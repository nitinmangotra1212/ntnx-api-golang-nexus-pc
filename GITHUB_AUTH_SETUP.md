# GitHub Authentication Setup for Maven Build

## Option 1: Use GitHub Personal Access Token (Recommended)

### Step 1: Create a GitHub Personal Access Token
1. Go to GitHub → Settings → Developer settings → Personal access tokens → Tokens (classic)
2. Click "Generate new token (classic)"
3. Give it a name (e.g., "Maven Build")
4. Select scopes: `repo` (full control of private repositories)
5. Click "Generate token"
6. **Copy the token** (you won't see it again!)

### Step 2: Configure Git to Use Token
```bash
# Configure Git credential helper
git config --global credential.helper store

# Test by cloning (will prompt for username and password/token)
# Username: your-github-username
# Password: <paste-your-token-here>
```

### Step 3: Update repositories.yaml to use HTTPS
The file should already have HTTPS URL. If not, ensure it's:
```yaml
- name: "common"
  type: "git"
  uri: "https://github.com/nutanix-core/ntnx-api-dev-platform.git"
  ref: "refs/tags/17.6.0.9581-RELEASE"
  baseDir: "ntnx-api-common/common-api-definitions/defs"
  authRequired: false
```

### Step 4: Trigger Git Credential Prompt
```bash
# This will prompt for credentials and store them
cd /tmp
git clone https://github.com/nutanix-core/ntnx-api-dev-platform.git test-clone
rm -rf test-clone
```

### Step 5: Run Maven Build
```bash
cd ~/ntnx-api-golang-mock-pc
mvn clean install -DskipTests -s settings.xml
```

---

## Option 2: Use SSH (If you have SSH keys set up)

### Step 1: Add SSH Key to GitHub
```bash
# Copy your public key
cat ~/.ssh/id_ed25519.pub
# Or
cat ~/.ssh/id_rsa.pub
```

1. Go to GitHub → Settings → SSH and GPG keys
2. Click "New SSH key"
3. Paste your public key
4. Save

### Step 2: Test SSH Connection
```bash
ssh -T git@github.com
# Should say: "Hi username! You've successfully authenticated..."
```

### Step 3: Update repositories.yaml to use SSH
```yaml
- name: "common"
  type: "git"
  uri: "git@github.com:nutanix-core/ntnx-api-dev-platform.git"
  ref: "refs/tags/17.6.0.9581-RELEASE"
  baseDir: "ntnx-api-common/common-api-definitions/defs"
  authRequired: false
```

### Step 4: Run Maven Build
```bash
cd ~/ntnx-api-golang-mock-pc
mvn clean install -DskipTests -s settings.xml
```

---

## Option 3: Workaround - Use Local Repository (Temporary)

If authentication setup is taking time, you can temporarily work around by:

1. The Maven plugin clones repos to a cache directory
2. You can manually copy your local repo to that cache location
3. Or modify the plugin to skip fetching if local copy exists

**Note**: This is a workaround and not recommended for production.

---

## Quick Test Commands

```bash
# Test HTTPS with token
git ls-remote https://github.com/nutanix-core/ntnx-api-dev-platform.git

# Test SSH
ssh -T git@github.com

# Check current Git config
git config --global --list | grep -i credential
git config --global --list | grep -i url
```

---

**Recommended**: Use Option 1 (Personal Access Token) as it's the most reliable for Maven builds.

