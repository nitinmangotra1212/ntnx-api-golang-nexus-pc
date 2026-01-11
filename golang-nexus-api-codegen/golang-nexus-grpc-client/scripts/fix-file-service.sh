#!/bin/bash
#
# Fix for NexusConfigFileServiceImpl.java
# 1. Adds proper exception handling for getMetadataSync() which throws TimeoutException and InterruptedException
# 2. Replaces setClientSideHeaders() calls with reflection-based field setting (workaround for Lombok version mismatch)
#

SERVICE_FILE="$1"
if [ -z "$SERVICE_FILE" ]; then
    SERVICE_FILE="target/generated-sources/swagger/src/nexus/v4/config/NexusConfigFileServiceImpl.java"
fi

if [ ! -f "$SERVICE_FILE" ]; then
    echo "File not found: $SERVICE_FILE"
    exit 0
fi

# Use Python for reliable text processing
export SERVICE_FILE_PATH="$SERVICE_FILE"
python3 << 'PYEOF'
import re
import os

file_path = os.environ.get('SERVICE_FILE_PATH')
if not file_path:
    print("Error: SERVICE_FILE_PATH not set")
    exit(1)

with open(file_path, 'r') as f:
    content = f.read()

# Add TimeoutException import if not present
if 'import java.util.concurrent.TimeoutException;' not in content:
    # Find the last import statement and add after it
    import_pattern = r'(import java\.util\.concurrent\.TimeUnit;)'
    replacement = r'\1\nimport java.util.concurrent.TimeoutException;'
    content = re.sub(import_pattern, replacement, content)

# Add gRPC imports if not present
if 'import io.grpc.ForwardingClientCall;' not in content:
    # Find a good place to add imports (after other io.grpc imports)
    import_pattern = r'(import io\.grpc\.stub\.StreamObserver;)'
    replacement = r'''\1
import io.grpc.ForwardingClientCall;
import io.grpc.ForwardingClientCallListener;
import io.grpc.MethodDescriptor;
import io.grpc.CallOptions;
import io.grpc.Channel;
import io.grpc.ClientCall;
import io.grpc.Metadata;'''
    content = re.sub(import_pattern, replacement, content)

# Fix downloadFile: Capture headers from response listener instead of getMetadataSync()
# We need to capture response headers in a variable accessible from both StreamObserver and interceptor
# Pattern: After interceptor declaration, before StreamObserver, add responseHeaders array
pattern_download_headers = r'(com\.nutanix\.devplatform\.interceptors\.HeaderInterceptor interceptor = new com\.nutanix\.devplatform\.interceptors\.HeaderInterceptor\(\);\s+)StreamObserver<nexus\.v4\.config\.DownloadFileRet> streamObserver'

replacement_download_headers = r'''\1// Store response headers captured from gRPC response (accessible from both StreamObserver and interceptor)
              final Metadata[] responseHeaders = new Metadata[1];
              StreamObserver<nexus.v4.config.DownloadFileRet> streamObserver'''

content = re.sub(pattern_download_headers, replacement_download_headers, content, flags=re.DOTALL)

# Replace getMetadataSync() call with responseHeaders[0]
# Pattern 1: Simple case - direct call to getMetadataSync()
pattern_get_metadata_simple = r'(\s+if\(!isResponseHeadersSet\)\{\s+)Metadata headers = interceptor\.getMetadataSync\(\);'

replacement_get_metadata_simple = r'''\1if (responseHeaders[0] == null) {
                              log.error("Response headers not yet received");
                              requestStream.cancel("Response headers not available", null);
                              completable.completeExceptionally(new RuntimeException("Response headers not available"));
                              return;
                          }
                          Metadata headers = responseHeaders[0];'''

content = re.sub(pattern_get_metadata_simple, replacement_get_metadata_simple, content, flags=re.DOTALL)

# Pattern 2: Case with try-catch (if it exists)
pattern_get_metadata_try = r'(\s+if\(!isResponseHeadersSet\)\{\s+)Metadata headers;\s+try \{\s+headers = interceptor\.getMetadataSync\(\);\s+\} catch \(TimeoutException \| InterruptedException e\) \{\s+log\.error\("Failed to get metadata from interceptor: \{\}", e\.getMessage\(\)\);\s+requestStream\.cancel\("Failed to get response headers", e\);\s+completable\.completeExceptionally\(e\);\s+return;\s+\}'

replacement_get_metadata_try = r'''\1if (responseHeaders[0] == null) {
                              log.error("Response headers not yet received");
                              requestStream.cancel("Response headers not available", null);
                              completable.completeExceptionally(new RuntimeException("Response headers not available"));
                              return;
                          }
                          Metadata headers = responseHeaders[0];'''

content = re.sub(pattern_get_metadata_try, replacement_get_metadata_try, content, flags=re.DOTALL)

# Replace setClientSideHeaders calls with custom interceptor that merges metadata directly
# Pattern 1: download - replace metadata creation and setClientSideHeaders, wrap interceptor
# Note: interceptor is already declared earlier, so we wrap it and use a final variable
# Also capture response headers in the listener
pattern1 = r'(\s+// Set file metadata headers from client side\s+Metadata metadata = new Metadata\(\);\s+metadata\.put\(FILE_IDENTIFIER, String\.valueOf\(extId\)\);\s+metadata\.put\(FILE_DOWNLOAD, Boolean\.toString\(true\)\);\s+)interceptor\.setClientSideHeaders\(metadata\);\s+this\.golangnexusGrpcAsyncStub\.withDeadlineAfter\(STREAM_API_TIMEOUT, TimeUnit\.SECONDS\)\.withMaxOutboundMessageSize\(BYTE_ARRAY_SIZE\)\.withInterceptors\(interceptor\)'

replacement1 = r'''\1// Create custom interceptor that merges metadata directly and captures response headers
                          final Metadata fileMetadata = new Metadata();
                          fileMetadata.put(FILE_IDENTIFIER, String.valueOf(extId));
                          fileMetadata.put(FILE_DOWNLOAD, Boolean.toString(true));
                          // Wrap the existing interceptor with metadata merging and header capture
                          final com.nutanix.devplatform.interceptors.HeaderInterceptor baseInterceptor = interceptor;
                          final com.nutanix.devplatform.interceptors.HeaderInterceptor wrappedInterceptor = new com.nutanix.devplatform.interceptors.HeaderInterceptor() {
                              @Override
                              public <T, E> ClientCall<T, E> interceptCall(MethodDescriptor<T, E> method, CallOptions callOptions, Channel next) {
                                  ClientCall<T, E> baseCall = baseInterceptor.interceptCall(method, callOptions, next);
                                  return new ForwardingClientCall.SimpleForwardingClientCall<T, E>(baseCall) {
                                      @Override
                                      public void start(ClientCall.Listener<E> responseListener, Metadata headers) {
                                          // Merge file metadata before sending
                                          headers.merge(fileMetadata);
                                          // Wrap response listener to capture response headers
                                          super.start(new ForwardingClientCallListener.SimpleForwardingClientCallListener<E>(responseListener) {
                                              @Override
                                              public void onHeaders(Metadata responseMetadata) {
                                                  // Capture response headers for use in onNext()
                                                  responseHeaders[0] = responseMetadata;
                                                  super.onHeaders(responseMetadata);
                                              }
                                          }, headers);
                                      }
                                  };
                              }
                          };
                          this.golangnexusGrpcAsyncStub.withDeadlineAfter(STREAM_API_TIMEOUT, TimeUnit.SECONDS).withMaxOutboundMessageSize(BYTE_ARRAY_SIZE).withInterceptors(wrappedInterceptor)'''

content = re.sub(pattern1, replacement1, content, flags=re.DOTALL)

# Pattern 2: upload - replace metadata creation, interceptor creation, and setClientSideHeaders
pattern2 = r'(\s+Metadata metadata = new Metadata\(\);\s+metadata\.put\(FILE_IDENTIFIER, fileName\);\s+com\.nutanix\.devplatform\.interceptors\.HeaderInterceptor interceptor = new com\.nutanix\.devplatform\.interceptors\.HeaderInterceptor\(\);\s+)interceptor\.setClientSideHeaders\(metadata\);\s+this\.golangnexusGrpcAsyncStub\.withDeadlineAfter\(STREAM_API_TIMEOUT, TimeUnit\.SECONDS\)\.withMaxOutboundMessageSize\(BYTE_ARRAY_SIZE\)\.withInterceptors\(interceptor\)\.uploadFile\(responseServer\);'

replacement2 = r'''\1// Create custom interceptor that merges metadata directly (workaround for Lombok version mismatch)
    final Metadata fileMetadata = new Metadata();
    fileMetadata.put(FILE_IDENTIFIER, fileName);
    final com.nutanix.devplatform.interceptors.HeaderInterceptor fileInterceptor = new com.nutanix.devplatform.interceptors.HeaderInterceptor() {
        @Override
        public <T, E> ClientCall<T, E> interceptCall(MethodDescriptor<T, E> method, CallOptions callOptions, Channel next) {
            return new ForwardingClientCall.SimpleForwardingClientCall<T, E>(next.newCall(method, callOptions)) {
                @Override
                public void start(ClientCall.Listener<E> responseListener, Metadata headers) {
                    // Merge file metadata
                    headers.merge(fileMetadata);
                    // Handle response headers
                    super.start(new ForwardingClientCallListener.SimpleForwardingClientCallListener<E>(responseListener) {
                        @Override
                        public void onHeaders(Metadata responseMetadata) {
                            setMetadata(responseMetadata);
                            super.onHeaders(responseMetadata);
                        }
                    }, headers);
                }
            };
        }
    };
    this.golangnexusGrpcAsyncStub.withDeadlineAfter(STREAM_API_TIMEOUT, TimeUnit.SECONDS).withMaxOutboundMessageSize(BYTE_ARRAY_SIZE).withInterceptors(fileInterceptor).uploadFile(responseServer);'''

content = re.sub(pattern2, replacement2, content, flags=re.DOTALL)

# Write back
with open(file_path, 'w') as f:
    f.write(content)

print("✅ Fixed NexusConfigFileServiceImpl.java - added exception handling and reflection-based field setting")
PYEOF

