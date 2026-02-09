# Security Policy

## Reporting a Vulnerability

The BanglaCode team takes security vulnerabilities seriously. We appreciate your efforts to responsibly disclose your findings.

### How to Report

**Please do NOT report security vulnerabilities through public GitHub issues.**

Instead, please report them via one of the following methods:

1. **GitHub Security Advisories** (Preferred)
   - Go to the [Security tab](https://github.com/nexoral/BanglaCode/security/advisories) of this repository
   - Click "New draft security advisory"
   - Fill in the details of the vulnerability

2. **Email**
   - Send an email to: [security@banglacode.dev](mailto:security@banglacode.dev)
   - Use the subject line: `[SECURITY] Brief description of vulnerability`

### What to Include

Please include the following information in your report:

- **Type of vulnerability** (e.g., code injection, path traversal, denial of service)
- **Location** — Full paths of source file(s) related to the vulnerability
- **Configuration** — Any special configuration required to reproduce the issue
- **Reproduction steps** — Step-by-step instructions to reproduce the issue
- **Proof of concept** — Code or commands that demonstrate the vulnerability
- **Impact** — What an attacker could achieve by exploiting this vulnerability
- **Suggested fix** — If you have ideas on how to fix the issue (optional)

### Response Timeline

| Action | Timeline |
|--------|----------|
| Initial acknowledgment | Within 48 hours |
| Initial assessment | Within 7 days |
| Status update | Every 14 days until resolved |
| Fix development | Depends on severity |
| Public disclosure | After fix is released |

### Severity Levels

We classify vulnerabilities using the following severity levels:

| Severity | Description | Response Time |
|----------|-------------|---------------|
| **Critical** | Remote code execution, complete system compromise | Immediate priority |
| **High** | Significant data exposure, privilege escalation | Within 7 days |
| **Medium** | Limited data exposure, denial of service | Within 30 days |
| **Low** | Minor issues, information disclosure | Within 90 days |

## Supported Versions

We provide security updates for the following versions:

| Version | Supported |
|---------|-----------|
| 3.x.x   | Yes |
| 2.x.x   | Security fixes only |
| < 2.0   | No |

We strongly recommend always using the latest version of BanglaCode.

## Security Considerations for BanglaCode

### Language Security Features

BanglaCode includes several security-conscious design decisions:

#### 1. File System Access

- File operations (`poro`, `lekho`) are restricted to the current directory and subdirectories by default
- Absolute paths outside the project directory require explicit configuration
- Symbolic link following can be disabled

#### 2. HTTP Server

- The built-in HTTP server binds to localhost by default
- External binding requires explicit configuration
- Basic input validation is performed on request data

#### 3. Module System

- Modules can only be imported from the local file system
- Remote module loading is not supported (by design)
- Circular import detection prevents infinite loops

#### 4. Code Execution

- No `eval()` function that executes arbitrary code strings
- No shell command execution built-in
- Memory limits prevent denial of service through excessive allocation

### Known Limitations

The following are known security limitations in the current version:

1. **No Sandboxing** — BanglaCode programs run with the same permissions as the interpreter
2. **File Access** — Programs can read/write files accessible to the user running the interpreter
3. **Network Access** — Programs can make outbound HTTP requests and bind to available ports

### Best Practices for Users

When running BanglaCode programs:

1. **Review code before execution** — Especially code from untrusted sources
2. **Use minimal permissions** — Run the interpreter with least-privilege user accounts
3. **Isolate execution** — Use containers or VMs for untrusted code
4. **Monitor network activity** — Be aware of programs using HTTP functions
5. **Limit file access** — Run from directories with only necessary files

### Best Practices for Contributors

When contributing to BanglaCode:

1. **Input validation** — Always validate input in built-in functions
2. **Error handling** — Never expose internal details in error messages
3. **Dependency management** — Keep Go dependencies updated
4. **Code review** — All changes require review before merging
5. **Testing** — Include tests for security-sensitive functionality

## Security Hardening Checklist

For production deployments:

- [ ] Use the latest stable version
- [ ] Run with minimal file system permissions
- [ ] Disable unnecessary built-in functions if possible
- [ ] Monitor resource usage (CPU, memory)
- [ ] Log program execution for audit purposes
- [ ] Use network isolation where appropriate
- [ ] Regularly update to receive security patches

## Vulnerability Disclosure Policy

### Our Commitment

- We will acknowledge receipt of your vulnerability report
- We will provide an estimated timeline for addressing the vulnerability
- We will notify you when the vulnerability is fixed
- We will publicly acknowledge your responsible disclosure (unless you prefer anonymity)

### Safe Harbor

We consider security research conducted in accordance with this policy to be:

- **Authorized** — We will not pursue legal action against researchers who follow this policy
- **Helpful** — We value your contribution to our security
- **Protected** — We will work with you to understand and resolve the issue

### Recognition

We maintain a [Security Hall of Fame](#security-hall-of-fame) to recognize individuals who have responsibly disclosed vulnerabilities.

## Security Hall of Fame

We thank the following individuals for responsibly disclosing security vulnerabilities:

*No submissions yet. Be the first to help improve BanglaCode's security!*

---

## Contact

For security-related questions that are not vulnerability reports:

- **GitHub Discussions**: [Security Category](https://github.com/nexoral/BanglaCode/discussions/categories/security)
- **Email**: [security@banglacode.dev](mailto:security@banglacode.dev)

---

**Last Updated**: 2024
**Policy Version**: 1.0
