# Getting Help with BanglaCode

Welcome! This document provides information on how to get help with BanglaCode.

## Quick Links

| Resource | Link | Purpose |
|----------|------|---------|
| Documentation | [banglacode.dev](https://banglacode.dev) | Complete language reference |
| GitHub Issues | [Issues](https://github.com/nexoral/BanglaCode/issues) | Bug reports, feature requests |
| GitHub Discussions | [Discussions](https://github.com/nexoral/BanglaCode/discussions) | Questions, ideas, community chat |
| Syntax Guide | [SYNTAX.md](SYNTAX.md) | Complete syntax reference |
| Examples | [examples/](examples/) | Sample programs |

---

## Before Asking for Help

Please try these steps first:

### 1. Check the Documentation

- **[README.md](README.md)** — Quick start and overview
- **[SYNTAX.md](SYNTAX.md)** — Complete language syntax
- **[Documentation Website](https://banglacode.dev)** — Full reference

### 2. Search Existing Resources

- **[GitHub Issues](https://github.com/nexoral/BanglaCode/issues)** — Someone may have had the same problem
- **[GitHub Discussions](https://github.com/nexoral/BanglaCode/discussions)** — Community Q&A
- **[Examples](examples/)** — Working code samples

### 3. Check Your Version

Make sure you're using the latest version:

```bash
./banglacode --version
```

Update if needed:

```bash
git pull origin main
go build -o banglacode main.go
```

---

## Types of Help

### Bug Reports

**Found a bug?** Please report it!

1. Go to [GitHub Issues](https://github.com/nexoral/BanglaCode/issues)
2. Click "New Issue"
3. Select "Bug Report"
4. Fill in the template with:
   - BanglaCode version
   - Operating system
   - Steps to reproduce
   - Expected vs actual behavior
   - Error messages (if any)

**Example Bug Report:**

```markdown
**Environment:**
- BanglaCode: v3.3.0
- OS: Ubuntu 22.04
- Go: 1.21.5

**Description:**
Array sorting fails for mixed types

**Steps to Reproduce:**
1. Create file with: `dhoro arr = [3, "a", 1]; saja(arr);`
2. Run: `./banglacode test.bang`
3. See error

**Expected:** Error message explaining mixed types can't be sorted
**Actual:** Program crashes without message

**Code:**
```banglacode
dhoro arr = [3, "a", 1];
saja(arr);
```
```

### Feature Requests

**Have an idea?** We'd love to hear it!

1. Go to [GitHub Issues](https://github.com/nexoral/BanglaCode/issues)
2. Click "New Issue"
3. Select "Feature Request"
4. Describe:
   - What you want to achieve
   - Proposed syntax (if applicable)
   - Use cases
   - Alternatives you've considered

### Questions

**Have a question?** Ask the community!

1. Go to [GitHub Discussions](https://github.com/nexoral/BanglaCode/discussions)
2. Click "New Discussion"
3. Select appropriate category:
   - **Q&A** — Technical questions
   - **Ideas** — Feature ideas and discussions
   - **Show and Tell** — Share your projects
   - **General** — Anything else

**Tips for Good Questions:**

- Use a clear, specific title
- Provide context and what you've tried
- Include minimal code examples
- Specify your BanglaCode version

---

## Common Issues

### Installation Problems

**Issue: `go build` fails**

```bash
# Make sure Go is installed
go version

# Should show Go 1.20 or higher
# If not, install/update Go from https://golang.org/dl/
```

**Issue: Command not found after building**

```bash
# Make sure you're in the right directory
cd /path/to/BanglaCode

# Build with explicit output
go build -o banglacode main.go

# Run with path
./banglacode
```

### Syntax Errors

**Issue: "unexpected token"**

Most common causes:
1. Missing semicolon (`;`)
2. Unbalanced braces (`{}`)
3. Typo in keyword

```banglacode
// Wrong
dhoro x = 5     // Missing semicolon

// Correct
dhoro x = 5;
```

**Issue: "variable not defined"**

```banglacode
// Wrong - typo in variable name
dhoro naam = "Ankan";
dekho(name);  // 'name' not 'naam'

// Correct
dekho(naam);
```

### Runtime Errors

**Issue: "index out of range"**

```banglacode
dhoro arr = [1, 2, 3];
dekho(arr[5]);  // Only indices 0, 1, 2 are valid

// Check length first
jodi (5 < dorghyo(arr)) {
    dekho(arr[5]);
}
```

**Issue: "division by zero"**

```banglacode
dhoro result = 10 / 0;  // Error!

// Check before dividing
jodi (divisor != 0) {
    dhoro result = 10 / divisor;
}
```

### VSCode Extension Issues

**Issue: Syntax highlighting not working**

1. Check file extension is `.bang`
2. Reload VSCode (Ctrl+Shift+P → "Reload Window")
3. Reinstall extension

**Issue: IntelliSense not working**

1. Ensure extension is enabled
2. Check for extension errors (View → Output → BanglaCode)
3. Update extension to latest version

---

## Getting Quick Answers

### REPL Help

Use the built-in help in the REPL:

```bash
./banglacode
>> sahajjo
```

This shows:
- All keywords
- Built-in functions
- Basic syntax examples

### Example Programs

Study the examples:

```bash
# List all examples
ls examples/

# Run an example
./banglacode examples/functions.bang
```

### Error Messages

BanglaCode error messages include:
- Line and column number
- Description of the error
- Context about what went wrong

```
Error [line 5, col 10]: variable 'x' is not defined
```

Use this information to locate and fix issues.

---

## Community Resources

### Learning Resources

| Resource | Type | Audience |
|----------|------|----------|
| [examples/hello.bang](examples/hello.bang) | Code | Beginners |
| [SYNTAX.md](SYNTAX.md) | Documentation | All levels |
| [examples/classes.bang](examples/classes.bang) | Code | Intermediate |
| [ARCHITECTURE.md](ARCHITECTURE.md) | Documentation | Advanced |

### Community Channels

- **GitHub Discussions** — Primary community forum
- **GitHub Issues** — Bug reports and feature requests

### Contributing

Want to help others? Consider:
- Answering questions in Discussions
- Improving documentation
- Creating tutorials
- Writing example programs

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

---

## Commercial Support

BanglaCode is open source and community-supported. For commercial inquiries or sponsorship:

- **Email**: [support@banglacode.dev](mailto:support@banglacode.dev)

---

## Response Times

| Channel | Expected Response |
|---------|-------------------|
| GitHub Issues (bugs) | 48-72 hours |
| GitHub Issues (features) | 1 week |
| GitHub Discussions | Community-driven |
| Security issues | 24-48 hours |

*Note: BanglaCode is maintained by volunteers. Response times may vary.*

---

## Helpful Tips

### Writing Good Code

1. **Use meaningful variable names** (in Banglish)
2. **Add comments** for complex logic
3. **Test incrementally** — don't write too much before testing
4. **Use the REPL** to experiment

### Debugging

1. **Use `dekho()`** to print values
2. **Check variable types** with `dhoron()`
3. **Simplify** — isolate the problem in minimal code
4. **Read error messages** — they tell you line and column

### Best Practices

```banglacode
// Good: descriptive names
dhoro studentNaam = "Rina";
dhoro studentBoyosh = 20;

// Good: comments for clarity
// Calculate total marks
kaj totalMarks(marks) {
    dhoro sum = 0;
    ghuriye (dhoro i = 0; i < dorghyo(marks); i = i + 1) {
        sum = sum + marks[i];
    }
    ferao sum;
}
```

---

## Still Need Help?

If you've tried the above and still need help:

1. **Create a minimal example** that reproduces your issue
2. **Post in GitHub Discussions** with:
   - What you're trying to do
   - What you've tried
   - The minimal code example
   - Error messages (if any)
   - Your environment (OS, BanglaCode version)

We're here to help! Don't hesitate to ask questions — that's how we all learn.

---

**সাহায্য প্রয়োজন? আমরা এখানে আছি!**

*Need help? We're here for you!*

---

**Last Updated**: 2024
