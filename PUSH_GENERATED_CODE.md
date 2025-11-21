# Pushing Generated-Code to Git Repository

## Overview

The `generated-code` directory contains all generated artifacts that are required for building the Go service. These files **must be committed and pushed** to the repository.

## What's in generated-code?

```
generated-code/
├── dto/                          # Go DTO definitions
│   └── src/
│       ├── go.mod
│       └── models/
│           └── mock/v4/
│               ├── config/
│               │   └── config_model.go
│               └── error/
│                   └── error_model.go
└── protobuf/                     # Protobuf definitions and generated Go code
    ├── mock/v4/config/           # Compiled protobuf Go files
    │   ├── cat_service.pb.go
    │   ├── cat_service_grpc.pb.go
    │   ├── config.pb.go
    │   └── go.mod
    └── swagger/mock/v4/          # Proto source files
        ├── api_version.proto
        ├── config/
        │   ├── cat_service.proto
        │   ├── config.proto
        │   └── config.pb.go
        ├── error/
        │   ├── error.proto
        │   └── error.pb.go
        └── http_method_options.proto
```

## Steps to Push Generated-Code

### Step 1: Verify Current Status

```bash
cd /Users/nitin.mangotra/ntnx-api-golang-mock-pc

# Check if generated-code is tracked
git ls-files generated-code/ | wc -l

# Check for uncommitted changes
git status generated-code/
```

### Step 2: Add Generated-Code Files

```bash
# Add all files in generated-code (including any modifications)
git add generated-code/

# Verify what will be committed
git status --short generated-code/
```

### Step 3: Commit Changes

```bash
# Commit with descriptive message
git commit -m "Update generated-code: DTOs, protobuf files, and Go modules

- Updated config_model.go with correct import paths
- Regenerated protobuf Go files
- Updated Go module files"
```

### Step 4: Push to Remote

```bash
# Push to remote repository
git push origin <your-branch>

# Or if pushing main branch
git push origin main
```

## Important Notes

1. **Always commit generated-code**: The Go service build depends on these files. Without them, the build will fail.

2. **After Maven build**: After running `mvn clean install`, always check if `generated-code` has new or modified files and commit them.

3. **Import path fixes**: If you manually fix import paths in generated files (like we did for `config_model.go`), ensure those fixes are committed.

4. **Go module files**: The `go.mod` files in `generated-code` are also important and should be committed.

## Verification

After pushing, verify the files are in the remote repository:

```bash
# Check remote status
git fetch origin
git log origin/<branch> --oneline -- generated-code/ | head -5

# Or check on GitHub/GitLab web interface
```

## Troubleshooting

### Issue: "nothing to commit, working tree clean"

**Solution**: This means all files in `generated-code` are already committed. You can verify with:
```bash
git log -1 --oneline -- generated-code/
```

### Issue: Files are ignored by .gitignore

**Solution**: Check `.gitignore` - `generated-code` should NOT be ignored. If it is, remove it from `.gitignore` and force add:
```bash
git add -f generated-code/
```

### Issue: Need to update after Maven build

**Solution**: After running `mvn clean install`, always check for new files:
```bash
# After Maven build
git status generated-code/
git add generated-code/
git commit -m "Update generated-code after Maven build"
git push origin <branch>
```

---

**Last Updated**: 2025-11-21

