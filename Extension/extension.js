const vscode = require('vscode');
const path = require('path');
const fs = require('fs');

/**
 * BanglaCode Language Extension
 * Provides IntelliSense and code completion for BanglaCode
 * Created by Ankan from West Bengal, India
 */

// All BanglaCode keywords
const keywords = [
    { label: 'dhoro', kind: vscode.CompletionItemKind.Keyword, detail: 'Variable declaration (ধরো)', insertText: 'dhoro ${1:name} = ${2:value};' },
    { label: 'jodi', kind: vscode.CompletionItemKind.Keyword, detail: 'If condition (যদি)', insertText: 'jodi (${1:condition}) {\n\t${2}\n}' },
    { label: 'nahole', kind: vscode.CompletionItemKind.Keyword, detail: 'Else (নাহলে)', insertText: 'nahole {\n\t${1}\n}' },
    { label: 'jotokkhon', kind: vscode.CompletionItemKind.Keyword, detail: 'While loop (যতক্ষণ)', insertText: 'jotokkhon (${1:condition}) {\n\t${2}\n}' },
    { label: 'ghuriye', kind: vscode.CompletionItemKind.Keyword, detail: 'For loop (ঘুরিয়ে)', insertText: 'ghuriye (dhoro ${1:i} = 0; ${1:i} < ${2:10}; ${1:i} = ${1:i} + 1) {\n\t${3}\n}' },
    { label: 'kaj', kind: vscode.CompletionItemKind.Keyword, detail: 'Function (কাজ)', insertText: 'kaj ${1:name}(${2:params}) {\n\t${3}\n}' },
    { label: 'ferao', kind: vscode.CompletionItemKind.Keyword, detail: 'Return (ফেরাও)', insertText: 'ferao ${1:value};' },
    { label: 'sreni', kind: vscode.CompletionItemKind.Keyword, detail: 'Class definition (শ্রেণী)', insertText: 'sreni ${1:Name} {\n\tshuru(${2:params}) {\n\t\t${3}\n\t}\n}' },
    { label: 'shuru', kind: vscode.CompletionItemKind.Keyword, detail: 'Constructor method (শুরু)', insertText: 'shuru(${1:params}) {\n\tei.${2:property} = ${2:property};\n}' },
    { label: 'notun', kind: vscode.CompletionItemKind.Keyword, detail: 'New instance (নতুন)', insertText: 'notun ${1:ClassName}(${2:args})' },
    { label: 'sotti', kind: vscode.CompletionItemKind.Constant, detail: 'True (সত্যি)', insertText: 'sotti' },
    { label: 'mittha', kind: vscode.CompletionItemKind.Constant, detail: 'False (মিথ্যা)', insertText: 'mittha' },
    { label: 'khali', kind: vscode.CompletionItemKind.Constant, detail: 'Null (খালি)', insertText: 'khali' },
    { label: 'ebong', kind: vscode.CompletionItemKind.Operator, detail: 'Logical AND (এবং)', insertText: 'ebong' },
    { label: 'ba', kind: vscode.CompletionItemKind.Operator, detail: 'Logical OR (বা)', insertText: 'ba' },
    { label: 'na', kind: vscode.CompletionItemKind.Operator, detail: 'Logical NOT (না)', insertText: 'na' },
    { label: 'thamo', kind: vscode.CompletionItemKind.Keyword, detail: 'Break (থামো)', insertText: 'thamo;' },
    { label: 'chharo', kind: vscode.CompletionItemKind.Keyword, detail: 'Continue (ছাড়ো)', insertText: 'chharo;' },
    { label: 'ei', kind: vscode.CompletionItemKind.Keyword, detail: 'This reference (এই)', insertText: 'ei.' },
    { label: 'ano', kind: vscode.CompletionItemKind.Keyword, detail: 'Import module (আনো)', insertText: 'ano "${1:module.bang}";' },
    { label: 'hisabe', kind: vscode.CompletionItemKind.Keyword, detail: 'Import alias (হিসাবে)', insertText: 'hisabe ${1:alias}' },
    { label: 'pathao', kind: vscode.CompletionItemKind.Keyword, detail: 'Export (পাঠাও)', insertText: 'pathao ' },
    { label: 'chesta', kind: vscode.CompletionItemKind.Keyword, detail: 'Try block (চেষ্টা)', insertText: 'chesta {\n\t${1}\n} dhoro_bhul (${2:err}) {\n\t${3}\n}' },
    { label: 'dhoro_bhul', kind: vscode.CompletionItemKind.Keyword, detail: 'Catch block (ধরো ভুল)', insertText: 'dhoro_bhul (${1:err}) {\n\t${2}\n}' },
    { label: 'shesh', kind: vscode.CompletionItemKind.Keyword, detail: 'Finally block (শেষ)', insertText: 'shesh {\n\t${1}\n}' },
    { label: 'felo', kind: vscode.CompletionItemKind.Keyword, detail: 'Throw error (ফেলো)', insertText: 'felo "${1:error message}";' },
];

// Built-in functions
const builtinFunctions = [
    // Output
    { label: 'dekho', kind: vscode.CompletionItemKind.Function, detail: 'দেখো - Print values', insertText: 'dekho(${1:value});', documentation: 'Prints values to the console' },
    
    // Type functions
    { label: 'dhoron', kind: vscode.CompletionItemKind.Function, detail: 'ধরন - Get type', insertText: 'dhoron(${1:value})', documentation: 'Returns the type of a value' },
    { label: 'lipi', kind: vscode.CompletionItemKind.Function, detail: 'লিপি - To string', insertText: 'lipi(${1:value})', documentation: 'Converts value to string' },
    { label: 'sonkha', kind: vscode.CompletionItemKind.Function, detail: 'সংখ্যা - To number', insertText: 'sonkha(${1:value})', documentation: 'Converts value to number' },
    { label: 'dorghyo', kind: vscode.CompletionItemKind.Function, detail: 'দৈর্ঘ্য - Get length', insertText: 'dorghyo(${1:value})', documentation: 'Returns the length of string or array' },
    
    // String functions
    { label: 'boroHater', kind: vscode.CompletionItemKind.Function, detail: 'বড় হাতের - Uppercase', insertText: 'boroHater(${1:str})', documentation: 'Converts string to uppercase' },
    { label: 'chotoHater', kind: vscode.CompletionItemKind.Function, detail: 'ছোট হাতের - Lowercase', insertText: 'chotoHater(${1:str})', documentation: 'Converts string to lowercase' },
    { label: 'chhanto', kind: vscode.CompletionItemKind.Function, detail: 'ছাঁটো - Trim', insertText: 'chhanto(${1:str})', documentation: 'Removes leading and trailing whitespace' },
    { label: 'bhag', kind: vscode.CompletionItemKind.Function, detail: 'ভাগ - Split string', insertText: 'bhag(${1:str}, ${2:separator})', documentation: 'Splits string into array' },
    { label: 'joro', kind: vscode.CompletionItemKind.Function, detail: 'জোড়ো - Join array', insertText: 'joro(${1:arr}, ${2:separator})', documentation: 'Joins array into string' },
    { label: 'khojo', kind: vscode.CompletionItemKind.Function, detail: 'খোঁজো - Index of', insertText: 'khojo(${1:str}, ${2:substr})', documentation: 'Finds index of substring (-1 if not found)' },
    { label: 'angsho', kind: vscode.CompletionItemKind.Function, detail: 'অংশ - Substring', insertText: 'angsho(${1:str}, ${2:start}, ${3:end})', documentation: 'Extracts substring from start to end' },
    { label: 'bodlo', kind: vscode.CompletionItemKind.Function, detail: 'বদলো - Replace', insertText: 'bodlo(${1:str}, ${2:old}, ${3:new})', documentation: 'Replaces all occurrences' },
    
    // Array functions
    { label: 'dhokao', kind: vscode.CompletionItemKind.Function, detail: 'ঢোকাও - Push to array', insertText: 'dhokao(${1:arr}, ${2:value})', documentation: 'Adds element to array' },
    { label: 'berKoro', kind: vscode.CompletionItemKind.Function, detail: 'বের করো - Pop from array', insertText: 'berKoro(${1:arr})', documentation: 'Removes and returns last element' },
    { label: 'kato', kind: vscode.CompletionItemKind.Function, detail: 'কাটো - Slice', insertText: 'kato(${1:arr}, ${2:start}, ${3:end})', documentation: 'Extracts portion of array' },
    { label: 'ulto', kind: vscode.CompletionItemKind.Function, detail: 'উল্টো - Reverse', insertText: 'ulto(${1:arr})', documentation: 'Reverses array (returns new array)' },
    { label: 'saja', kind: vscode.CompletionItemKind.Function, detail: 'সাজা - Sort', insertText: 'saja(${1:arr})', documentation: 'Sorts array (returns new array)' },
    { label: 'ache', kind: vscode.CompletionItemKind.Function, detail: 'আছে - Includes', insertText: 'ache(${1:arr}, ${2:value})', documentation: 'Checks if value exists in array' },
    { label: 'chabi', kind: vscode.CompletionItemKind.Function, detail: 'চাবি - Get keys', insertText: 'chabi(${1:map})', documentation: 'Returns array of map keys' },
    
    // Math functions
    { label: 'borgomul', kind: vscode.CompletionItemKind.Function, detail: 'বর্গমূল - Square root', insertText: 'borgomul(${1:x})', documentation: 'Returns square root' },
    { label: 'ghat', kind: vscode.CompletionItemKind.Function, detail: 'ঘাত - Power', insertText: 'ghat(${1:base}, ${2:exp})', documentation: 'Returns base raised to exponent' },
    { label: 'niche', kind: vscode.CompletionItemKind.Function, detail: 'নিচে - Floor', insertText: 'niche(${1:x})', documentation: 'Rounds down to integer' },
    { label: 'upore', kind: vscode.CompletionItemKind.Function, detail: 'উপরে - Ceiling', insertText: 'upore(${1:x})', documentation: 'Rounds up to integer' },
    { label: 'kache', kind: vscode.CompletionItemKind.Function, detail: 'কাছে - Round', insertText: 'kache(${1:x})', documentation: 'Rounds to nearest integer' },
    { label: 'niratek', kind: vscode.CompletionItemKind.Function, detail: 'নিরপেক্ষ - Absolute', insertText: 'niratek(${1:x})', documentation: 'Returns absolute value' },
    { label: 'choto', kind: vscode.CompletionItemKind.Function, detail: 'ছোট - Minimum', insertText: 'choto(${1:a}, ${2:b})', documentation: 'Returns minimum value' },
    { label: 'boro', kind: vscode.CompletionItemKind.Function, detail: 'বড় - Maximum', insertText: 'boro(${1:a}, ${2:b})', documentation: 'Returns maximum value' },
    { label: 'lotto', kind: vscode.CompletionItemKind.Function, detail: 'লটো - Random', insertText: 'lotto()', documentation: 'Returns random number between 0 and 1' },
    
    // Utility functions
    { label: 'somoy', kind: vscode.CompletionItemKind.Function, detail: 'সময় - Current time', insertText: 'somoy()', documentation: 'Returns current timestamp in milliseconds' },
    { label: 'ghum', kind: vscode.CompletionItemKind.Function, detail: 'ঘুম - Sleep', insertText: 'ghum(${1:milliseconds})', documentation: 'Pauses execution for milliseconds' },
    { label: 'nao', kind: vscode.CompletionItemKind.Function, detail: 'নাও - Input', insertText: 'nao(${1:"Enter value: "})', documentation: 'Reads user input from console' },
    { label: 'bondho', kind: vscode.CompletionItemKind.Function, detail: 'বন্ধ - Exit', insertText: 'bondho(${1:0})', documentation: 'Exits program with code' },
    
    // File functions
    { label: 'poro', kind: vscode.CompletionItemKind.Function, detail: 'পড়ো - Read file', insertText: 'poro(${1:"filename"})', documentation: 'Reads file contents as string' },
    { label: 'lekho', kind: vscode.CompletionItemKind.Function, detail: 'লেখো - Write file', insertText: 'lekho(${1:"filename"}, ${2:content})', documentation: 'Writes content to file' },
    
    // HTTP functions
    { label: 'server_chalu', kind: vscode.CompletionItemKind.Function, detail: 'সার্ভার চালু - Start HTTP server', insertText: 'server_chalu(${1:3000}, ${2:handler})', documentation: 'Starts HTTP server on port' },
    { label: 'anun', kind: vscode.CompletionItemKind.Function, detail: 'আনুন - HTTP GET', insertText: 'anun(${1:"url"})', documentation: 'Makes HTTP GET request' },
];

/**
 * Extract @comment documentation for a function or class
 * Looks for // @comment lines above the definition
 * @param {string} text - Full document text
 * @param {string} name - Name of function/class/method
 * @param {string} type - 'kaj', 'sreni', or 'shuru'
 * @returns {string|null} - Documentation text or null
 */
function extractDocComment(text, name, type = 'kaj') {
    // Build regex to find the definition
    let defRegex;
    if (type === 'sreni') {
        defRegex = new RegExp(`((?:\\/\\/[^\\n]*\\n)*)\\s*(?:pathao\\s+)?sreni\\s+${name}\\s*\\{`, 'm');
    } else if (type === 'shuru') {
        defRegex = new RegExp(`((?:\\/\\/[^\\n]*\\n)*)\\s*shuru\\s*\\(`, 'm');
    } else {
        defRegex = new RegExp(`((?:\\/\\/[^\\n]*\\n)*)\\s*(?:pathao\\s+)?kaj\\s+${name}\\s*\\(`, 'm');
    }
    
    const match = defRegex.exec(text);
    if (!match || !match[1]) return null;
    
    // Extract @comment lines from the comment block
    const commentBlock = match[1];
    const lines = commentBlock.split('\n').filter(l => l.trim());
    const docLines = [];
    
    for (const line of lines) {
        const commentMatch = line.match(/\/\/\s*@comment\s+(.*)/);
        if (commentMatch) {
            docLines.push(commentMatch[1].trim());
        }
    }
    
    // Format each line as a separate item for proper markdown display
    return docLines.length > 0 ? docLines.map(l => `• ${l}`).join('\n\n') : null;
}

/**
 * Extract @comment docs from an imported module file
 */
function extractDocFromModule(modulePath, name, currentDocPath) {
    try {
        const currentDir = path.dirname(currentDocPath);
        const fullPath = path.resolve(currentDir, modulePath);
        
        if (!fs.existsSync(fullPath)) return null;
        
        const moduleText = fs.readFileSync(fullPath, 'utf8');
        return extractDocComment(moduleText, name, 'kaj');
    } catch (err) {
        return null;
    }
}

/**
 * Extract map/JSON keys from a variable declaration
 * Supports nested maps and detects value types
 */
function extractMapKeys(text, varName) {
    const keys = [];
    
    // Find the map declaration: dhoro varName = { ... }
    // Use a simple brace-matching approach
    const startRegex = new RegExp(`dhoro\\s+${varName}\\s*=\\s*\\{`);
    const startMatch = startRegex.exec(text);
    if (!startMatch) return keys;
    
    const startIndex = startMatch.index + startMatch[0].length - 1; // Position of opening brace
    let braceCount = 0;
    let endIndex = startIndex;
    
    // Find matching closing brace
    for (let i = startIndex; i < text.length; i++) {
        if (text[i] === '{') braceCount++;
        else if (text[i] === '}') {
            braceCount--;
            if (braceCount === 0) {
                endIndex = i;
                break;
            }
        }
    }
    
    const mapContent = text.substring(startIndex + 1, endIndex);
    
    // Parse top-level keys (simple regex approach for common cases)
    // Match: key: value (where key is identifier)
    const keyRegex = /^\s*([a-zA-Z_][a-zA-Z0-9_]*)\s*:\s*/gm;
    let keyMatch;
    
    while ((keyMatch = keyRegex.exec(mapContent)) !== null) {
        const keyName = keyMatch[1];
        const valueStart = keyMatch.index + keyMatch[0].length;
        
        // Determine value type
        let valueType = 'unknown';
        let nestedKeys = null;
        const restOfLine = mapContent.substring(valueStart);
        
        if (restOfLine.startsWith('"') || restOfLine.startsWith("'")) {
            valueType = 'string';
        } else if (restOfLine.startsWith('[')) {
            valueType = 'array';
        } else if (restOfLine.startsWith('{')) {
            valueType = 'object';
            // Extract nested keys
            nestedKeys = [];
            let nestedBraceCount = 0;
            let nestedEndIndex = 0;
            for (let i = 0; i < restOfLine.length; i++) {
                if (restOfLine[i] === '{') nestedBraceCount++;
                else if (restOfLine[i] === '}') {
                    nestedBraceCount--;
                    if (nestedBraceCount === 0) {
                        nestedEndIndex = i;
                        break;
                    }
                }
            }
            const nestedContent = restOfLine.substring(1, nestedEndIndex);
            const nestedKeyRegex = /([a-zA-Z_][a-zA-Z0-9_]*)\s*:/g;
            let nestedMatch;
            while ((nestedMatch = nestedKeyRegex.exec(nestedContent)) !== null) {
                nestedKeys.push(nestedMatch[1]);
            }
        } else if (/^[0-9]/.test(restOfLine) || /^-[0-9]/.test(restOfLine)) {
            valueType = 'number';
        } else if (restOfLine.startsWith('sotti') || restOfLine.startsWith('mittha')) {
            valueType = 'boolean';
        }
        
        keys.push({ name: keyName, type: valueType, nested: nestedKeys });
    }
    
    return keys;
}

/**
 * Extract keys from a parsed JSON object
 */
function extractJSONKeys(jsonData, prefix = '') {
    const keys = [];
    
    if (typeof jsonData !== 'object' || jsonData === null) {
        return keys;
    }
    
    for (const [key, value] of Object.entries(jsonData)) {
        let valueType = 'unknown';
        let nestedKeys = null;
        
        if (value === null) {
            valueType = 'null';
        } else if (typeof value === 'string') {
            valueType = 'string';
        } else if (typeof value === 'number') {
            valueType = 'number';
        } else if (typeof value === 'boolean') {
            valueType = 'boolean';
        } else if (Array.isArray(value)) {
            valueType = 'array';
        } else if (typeof value === 'object') {
            valueType = 'object';
            nestedKeys = Object.keys(value);
        }
        
        keys.push({ name: key, type: valueType, nested: nestedKeys });
    }
    
    return keys;
}

function activate(context) {
    console.log('BanglaCode extension activated!');

    // Register completion provider
    const completionProvider = vscode.languages.registerCompletionItemProvider(
        'banglacode',
        {
            provideCompletionItems(document, position, token, context) {
                const completions = [];
                
                // Add keywords
                for (const kw of keywords) {
                    const item = new vscode.CompletionItem(kw.label, kw.kind);
                    item.detail = kw.detail;
                    item.insertText = new vscode.SnippetString(kw.insertText);
                    completions.push(item);
                }
                
                // Add built-in functions
                for (const fn of builtinFunctions) {
                    const item = new vscode.CompletionItem(fn.label, fn.kind);
                    item.detail = fn.detail;
                    item.insertText = new vscode.SnippetString(fn.insertText);
                    if (fn.documentation) {
                        item.documentation = new vscode.MarkdownString(fn.documentation);
                    }
                    completions.push(item);
                }
                
                // Extract variables from current document
                const text = document.getText();
                const varRegex = /dhoro\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*=/g;
                let match;
                const seenVars = new Set();
                while ((match = varRegex.exec(text)) !== null) {
                    const varName = match[1];
                    if (!seenVars.has(varName)) {
                        seenVars.add(varName);
                        const item = new vscode.CompletionItem(varName, vscode.CompletionItemKind.Variable);
                        item.detail = 'Variable';
                        completions.push(item);
                    }
                }
                
                // Extract functions from current document
                const funcRegex = /kaj\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(/g;
                const seenFuncs = new Set();
                while ((match = funcRegex.exec(text)) !== null) {
                    const funcName = match[1];
                    if (!seenFuncs.has(funcName) && funcName !== 'shuru') {
                        seenFuncs.add(funcName);
                        const item = new vscode.CompletionItem(funcName, vscode.CompletionItemKind.Function);
                        item.detail = 'User function';
                        item.insertText = new vscode.SnippetString(funcName + '(${1})');
                        completions.push(item);
                    }
                }
                
                // Extract classes from current document
                const classRegex = /sreni\s+([a-zA-Z_][a-zA-Z0-9_]*)/g;
                const seenClasses = new Set();
                while ((match = classRegex.exec(text)) !== null) {
                    const className = match[1];
                    if (!seenClasses.has(className)) {
                        seenClasses.add(className);
                        const item = new vscode.CompletionItem(className, vscode.CompletionItemKind.Class);
                        item.detail = 'User class';
                        completions.push(item);
                    }
                }
                
                return completions;
            }
        },
        '' // Trigger on any character
    );

    // Register import path completion provider
    const importPathProvider = vscode.languages.registerCompletionItemProvider(
        'banglacode',
        {
            provideCompletionItems(document, position, token, context) {
                const line = document.lineAt(position.line).text;
                const linePrefix = line.substring(0, position.character);
                
                // Check if we're inside an import statement string
                const importMatch = linePrefix.match(/ano\s+["']([^"']*)$/);
                if (!importMatch) {
                    return undefined;
                }
                
                const typedPath = importMatch[1];
                const completions = [];
                
                // Get current file's directory
                const currentDir = path.dirname(document.uri.fsPath);
                
                // Determine the search directory based on typed path
                let searchDir = currentDir;
                let prefix = '';
                
                if (typedPath.includes('/')) {
                    const lastSlash = typedPath.lastIndexOf('/');
                    prefix = typedPath.substring(0, lastSlash + 1);
                    const relativePath = typedPath.substring(0, lastSlash);
                    searchDir = path.resolve(currentDir, relativePath);
                } else if (typedPath.startsWith('.')) {
                    // Just started typing ./ or ../
                    prefix = '';
                    searchDir = currentDir;
                }
                
                try {
                    if (!fs.existsSync(searchDir)) {
                        return completions;
                    }
                    
                    const entries = fs.readdirSync(searchDir, { withFileTypes: true });
                    
                    for (const entry of entries) {
                        // Skip hidden files
                        if (entry.name.startsWith('.')) continue;
                        
                        if (entry.isDirectory()) {
                            const item = new vscode.CompletionItem(
                                entry.name + '/',
                                vscode.CompletionItemKind.Folder
                            );
                            item.detail = 'Directory';
                            item.insertText = entry.name + '/';
                            item.command = {
                                command: 'editor.action.triggerSuggest',
                                title: 'Re-trigger completions'
                            };
                            completions.push(item);
                        } else if (entry.name.endsWith('.bang')) {
                            const item = new vscode.CompletionItem(
                                entry.name,
                                vscode.CompletionItemKind.File
                            );
                            item.detail = 'BanglaCode module';
                            item.insertText = entry.name;
                            completions.push(item);
                        } else if (entry.name.endsWith('.json')) {
                            const item = new vscode.CompletionItem(
                                entry.name,
                                vscode.CompletionItemKind.File
                            );
                            item.detail = 'JSON file';
                            item.insertText = entry.name;
                            completions.push(item);
                        }
                    }
                } catch (err) {
                    console.error('Error reading directory:', err);
                }
                
                return completions;
            }
        },
        '/', '.', '"', "'" // Trigger on these characters
    );

    // Register map/JSON property completion provider
    const mapPropertyProvider = vscode.languages.registerCompletionItemProvider(
        'banglacode',
        {
            provideCompletionItems(document, position, token, context) {
                const line = document.lineAt(position.line).text;
                const linePrefix = line.substring(0, position.character);
                
                // Check if we're after a dot (e.g., "json.")
                const dotMatch = linePrefix.match(/([a-zA-Z_][a-zA-Z0-9_]*)\.$/);
                if (!dotMatch) {
                    return undefined;
                }
                
                const varName = dotMatch[1];
                const text = document.getText();
                const completions = [];
                
                // First check if this is an import alias
                const aliasImportRegex = new RegExp(`ano\\s+"([^"]+)"\\s+hisabe\\s+${varName}\\s*;?`);
                const aliasImport = aliasImportRegex.exec(text);
                
                if (aliasImport) {
                    const modulePath = aliasImport[1];
                    try {
                        const currentDir = path.dirname(document.uri.fsPath);
                        const fullPath = path.resolve(currentDir, modulePath);
                        if (fs.existsSync(fullPath)) {
                            // Check if it's a JSON file
                            if (modulePath.endsWith('.json')) {
                                const jsonText = fs.readFileSync(fullPath, 'utf8');
                                try {
                                    const jsonData = JSON.parse(jsonText);
                                    const keys = extractJSONKeys(jsonData);
                                    for (const key of keys) {
                                        const item = new vscode.CompletionItem(key.name, vscode.CompletionItemKind.Property);
                                        item.detail = `${key.type} property`;
                                        if (key.nested) {
                                            item.documentation = new vscode.MarkdownString('Nested object with: ' + key.nested.join(', '));
                                        }
                                        completions.push(item);
                                    }
                                } catch (err) {
                                    console.error('Error parsing JSON:', err);
                                }
                            } else {
                                // It's a .bang module - get functions
                                const moduleText = fs.readFileSync(fullPath, 'utf8');
                                
                                // Find exported functions
                                const funcRegex = /pathao\s+kaj\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)/g;
                                let funcMatch;
                                while ((funcMatch = funcRegex.exec(moduleText)) !== null) {
                                    const funcName = funcMatch[1];
                                    const params = funcMatch[2].trim();
                                    const docComment = extractDocComment(moduleText, funcName, 'kaj');
                                    
                                    const item = new vscode.CompletionItem(funcName, vscode.CompletionItemKind.Method);
                                    item.detail = `${varName}.${funcName}(${params})`;
                                    item.insertText = new vscode.SnippetString(funcName + '(${1})');
                                    if (docComment) {
                                        item.documentation = new vscode.MarkdownString(docComment);
                                    }
                                    completions.push(item);
                                }
                            }
                        }
                    } catch (err) {
                        console.error('Error reading module:', err);
                    }
                    
                    return completions;
                }
                
                // Check if it's a map/JSON variable - extract its keys
                const mapKeys = extractMapKeys(text, varName);
                if (mapKeys.length > 0) {
                    for (const key of mapKeys) {
                        const item = new vscode.CompletionItem(key.name, vscode.CompletionItemKind.Property);
                        item.detail = key.type ? `Property (${key.type})` : 'Property';
                        if (key.nested) {
                            item.documentation = new vscode.MarkdownString('Nested object with properties: ' + key.nested.join(', '));
                        }
                        completions.push(item);
                    }
                    return completions;
                }
                
                return undefined;
            }
        },
        '.' // Trigger on dot
    );

    // Register hover provider
    const hoverProvider = vscode.languages.registerHoverProvider('banglacode', {
        provideHover(document, position, token) {
            const range = document.getWordRangeAtPosition(position);
            if (!range) return null;
            
            const word = document.getText(range);
            const text = document.getText();
            const line = document.lineAt(position.line).text;
            
            // Check keywords
            const kw = keywords.find(k => k.label === word);
            if (kw) {
                return new vscode.Hover([
                    `**${kw.label}** - ${kw.detail}`,
                    'BanglaCode keyword'
                ]);
            }
            
            // Check built-in functions
            const fn = builtinFunctions.find(f => f.label === word);
            if (fn) {
                return new vscode.Hover([
                    `**${fn.label}** - ${fn.detail}`,
                    fn.documentation || '',
                    'BanglaCode built-in function'
                ]);
            }
            
            // Check user-defined functions
            const funcRegex = new RegExp(`(?:pathao\\s+)?kaj\\s+${word}\\s*\\(([^)]*)\\)`, 'g');
            let funcMatch = funcRegex.exec(text);
            if (funcMatch) {
                const params = funcMatch[1].trim();
                const paramList = params ? params.split(',').map(p => p.trim()) : [];
                const paramCount = paramList.length;
                
                // Extract @comment documentation
                const docComment = extractDocComment(text, word, 'kaj');
                
                const md = new vscode.MarkdownString();
                md.appendMarkdown(`**kaj ${word}(${params})**\n\n`);
                
                if (docComment) {
                    md.appendMarkdown('---\n\n');
                    md.appendMarkdown(docComment + '\n\n');
                    md.appendMarkdown('---\n\n');
                }
                
                md.appendMarkdown(`**Parameters:** ${paramCount === 0 ? 'none' : paramList.join(', ')}\n\n`);
                md.appendMarkdown(`**Call with:** \`${word}(${paramList.map((_, i) => `arg${i+1}`).join(', ')})\``);
                
                return new vscode.Hover(md);
            }
            
            // Check import aliases - hover on alias name
            const aliasRegex = new RegExp(`ano\\s+"([^"]+)"\\s+hisabe\\s+${word}\\s*;?`);
            const aliasMatch = aliasRegex.exec(text);
            if (aliasMatch) {
                const modulePath = aliasMatch[1];
                
                // Check if it's a JSON file
                if (modulePath.endsWith('.json')) {
                    try {
                        const currentDir = path.dirname(document.uri.fsPath);
                        const fullPath = path.resolve(currentDir, modulePath);
                        if (fs.existsSync(fullPath)) {
                            const jsonText = fs.readFileSync(fullPath, 'utf8');
                            const jsonData = JSON.parse(jsonText);
                            const keys = Object.keys(jsonData);
                            
                            const md = new vscode.MarkdownString();
                            md.appendMarkdown(`**${word}** - Imported JSON\n\n`);
                            md.appendMarkdown(`📦 File: \`${modulePath}\`\n\n`);
                            md.appendMarkdown('---\n\n');
                            md.appendMarkdown('**Properties:**\n\n');
                            for (const key of keys.slice(0, 10)) {
                                const value = jsonData[key];
                                const type = Array.isArray(value) ? 'array' : typeof value;
                                md.appendMarkdown(`• \`${key}\`: ${type}\n\n`);
                            }
                            if (keys.length > 10) {
                                md.appendMarkdown(`... and ${keys.length - 10} more\n\n`);
                            }
                            md.appendMarkdown('---\n\n');
                            md.appendMarkdown(`**Access:** \`${word}.propertyName\``);
                            
                            return new vscode.Hover(md);
                        }
                    } catch (err) {
                        // Fallback to basic hover
                    }
                }
                
                // Regular .bang module
                return new vscode.Hover([
                    `**${word}** - Imported module alias`,
                    `Module: \`${modulePath}\``,
                    `Access with: ${word}.functionName(args)`
                ]);
            }
            
            // Check classes
            const classRegex = new RegExp(`(?:pathao\\s+)?sreni\\s+${word}\\s*\\{`);
            if (classRegex.test(text)) {
                // Find constructor parameters
                const classBodyRegex = new RegExp(`sreni\\s+${word}\\s*\\{([\\s\\S]*?)\\n\\}`, 'm');
                const classBody = classBodyRegex.exec(text);
                let constructorParams = '';
                if (classBody) {
                    const shuruMatch = classBody[1].match(/shuru\s*\(([^)]*)\)/);
                    if (shuruMatch) {
                        constructorParams = shuruMatch[1];
                    }
                }
                
                // Extract @comment documentation for class
                const docComment = extractDocComment(text, word, 'sreni');
                
                const md = new vscode.MarkdownString();
                md.appendMarkdown(`**sreni ${word}**\n\n`);
                
                if (docComment) {
                    md.appendMarkdown('---\n\n');
                    md.appendMarkdown(docComment + '\n\n');
                    md.appendMarkdown('---\n\n');
                }
                
                md.appendMarkdown(constructorParams ? `**Constructor:** shuru(${constructorParams})\n\n` : 'No constructor\n\n');
                md.appendMarkdown(`**Create:** \`notun ${word}(${constructorParams})\``);
                
                return new vscode.Hover(md);
            }
            
            // Check variables
            const varRegex = new RegExp(`dhoro\\s+${word}\\s*=`);
            if (varRegex.test(text)) {
                // Check if it's a map/JSON variable
                const mapKeys = extractMapKeys(text, word);
                if (mapKeys.length > 0) {
                    const md = new vscode.MarkdownString();
                    md.appendMarkdown(`**${word}** - Map/JSON Object\n\n`);
                    md.appendMarkdown('---\n\n');
                    md.appendMarkdown('**Properties:**\n\n');
                    for (const key of mapKeys) {
                        if (key.nested) {
                            md.appendMarkdown(`• \`${key.name}\`: object { ${key.nested.join(', ')} }\n\n`);
                        } else {
                            md.appendMarkdown(`• \`${key.name}\`: ${key.type}\n\n`);
                        }
                    }
                    md.appendMarkdown('---\n\n');
                    md.appendMarkdown(`**Access:** \`${word}.propertyName\``);
                    return new vscode.Hover(md);
                }
                
                return new vscode.Hover([
                    `**${word}** - Variable`,
                    'Declared with dhoro'
                ]);
            }
            
            // Check if hovering over a map property (e.g., json.name)
            const lineText = document.lineAt(position.line).text;
            const wordStart = position.character - word.length;
            const beforeWord = lineText.substring(0, wordStart);
            const mapPropertyMatch = beforeWord.match(/([a-zA-Z_][a-zA-Z0-9_]*)\.$/);
            
            if (mapPropertyMatch) {
                const mapName = mapPropertyMatch[1];
                
                // First check if it's an import alias
                const aliasImportRegex = new RegExp(`ano\\s+"([^"]+)"\\s+hisabe\\s+${mapName}\\s*;?`);
                const aliasImport = aliasImportRegex.exec(text);
                
                if (aliasImport) {
                    const modulePath = aliasImport[1];
                    try {
                        const currentDir = path.dirname(document.uri.fsPath);
                        const fullPath = path.resolve(currentDir, modulePath);
                        if (fs.existsSync(fullPath)) {
                            // Check if it's a JSON file
                            if (modulePath.endsWith('.json')) {
                                const jsonText = fs.readFileSync(fullPath, 'utf8');
                                try {
                                    const jsonData = JSON.parse(jsonText);
                                    if (jsonData.hasOwnProperty(word)) {
                                        const value = jsonData[word];
                                        const md = new vscode.MarkdownString();
                                        md.appendMarkdown(`**${mapName}.${word}**\n\n`);
                                        md.appendMarkdown(`📦 From JSON: \`${modulePath}\`\n\n`);
                                        
                                        if (typeof value === 'object' && value !== null && !Array.isArray(value)) {
                                            md.appendMarkdown(`**Type:** object\n\n`);
                                            md.appendMarkdown(`**Properties:** ${Object.keys(value).join(', ')}\n\n`);
                                            md.appendMarkdown(`**Access:** \`${mapName}.${word}.propertyName\``);
                                        } else if (Array.isArray(value)) {
                                            md.appendMarkdown(`**Type:** array\n\n`);
                                            md.appendMarkdown(`**Length:** ${value.length}`);
                                        } else {
                                            md.appendMarkdown(`**Type:** ${typeof value}\n\n`);
                                            md.appendMarkdown(`**Value:** \`${JSON.stringify(value)}\``);
                                        }
                                        
                                        return new vscode.Hover(md);
                                    }
                                } catch (err) {
                                    // Ignore JSON parse errors
                                }
                            } else {
                                // It's a .bang module
                                const moduleText = fs.readFileSync(fullPath, 'utf8');
                                
                                // Get function signature
                                const funcInModuleRegex = new RegExp(`pathao\\s+kaj\\s+${word}\\s*\\(([^)]*)\\)`);
                                const funcInModule = funcInModuleRegex.exec(moduleText);
                                
                                if (funcInModule) {
                                    const params = funcInModule[1].trim();
                                    const docComment = extractDocComment(moduleText, word, 'kaj');
                                    
                                    const md = new vscode.MarkdownString();
                                    md.appendMarkdown(`**${mapName}.${word}(${params})**\n\n`);
                                    md.appendMarkdown(`📦 Imported from: \`${modulePath}\`\n\n`);
                                    
                                    if (docComment) {
                                        md.appendMarkdown('---\n\n');
                                        md.appendMarkdown(docComment + '\n\n');
                                        md.appendMarkdown('---\n\n');
                                    }
                                    
                                    md.appendMarkdown(`**Parameters:** ${params || 'none'}\n\n`);
                                    md.appendMarkdown(`**Call:** \`${mapName}.${word}(${params.split(',').map((_, i) => `arg${i+1}`).join(', ')})\``);
                                    
                                    return new vscode.Hover(md);
                                }
                            }
                        }
                    } catch (err) {
                        // Ignore errors
                    }
                } else {
                    // Check if it's a map/JSON variable property
                    const mapKeys = extractMapKeys(text, mapName);
                    const keyInfo = mapKeys.find(k => k.name === word);
                    if (keyInfo) {
                        const md = new vscode.MarkdownString();
                        md.appendMarkdown(`**${mapName}.${word}**\n\n`);
                        if (keyInfo.nested) {
                            md.appendMarkdown(`**Type:** object\n\n`);
                            md.appendMarkdown(`**Nested properties:** ${keyInfo.nested.join(', ')}\n\n`);
                            md.appendMarkdown(`**Access nested:** \`${mapName}.${word}.propertyName\``);
                        } else {
                            md.appendMarkdown(`**Type:** ${keyInfo.type}\n\n`);
                            md.appendMarkdown(`Property of \`${mapName}\` object`);
                        }
                        return new vscode.Hover(md);
                    }
                }
            }
            
            // Check if function is from an imported module (direct import, no alias)
            const importRegex = /ano\s+"([^"]+)"\s*;/g;
            let importMatch;
            while ((importMatch = importRegex.exec(text)) !== null) {
                const modulePath = importMatch[1];
                const docComment = extractDocFromModule(modulePath, word, document.uri.fsPath);
                if (docComment) {
                    // Try to get function signature from module
                    try {
                        const currentDir = path.dirname(document.uri.fsPath);
                        const fullPath = path.resolve(currentDir, modulePath);
                        if (fs.existsSync(fullPath)) {
                            const moduleText = fs.readFileSync(fullPath, 'utf8');
                            const funcInModuleRegex = new RegExp(`pathao\\s+kaj\\s+${word}\\s*\\(([^)]*)\\)`);
                            const funcInModule = funcInModuleRegex.exec(moduleText);
                            if (funcInModule) {
                                const params = funcInModule[1].trim();
                                const md = new vscode.MarkdownString();
                                md.appendMarkdown(`**kaj ${word}(${params})**\n\n`);
                                md.appendMarkdown(`📦 Imported from: \`${modulePath}\`\n\n`);
                                md.appendMarkdown('---\n\n');
                                md.appendMarkdown(docComment + '\n\n');
                                md.appendMarkdown('---\n\n');
                                md.appendMarkdown(`**Parameters:** ${params || 'none'}`);
                                return new vscode.Hover(md);
                            }
                        }
                    } catch (err) {
                        // Ignore errors
                    }
                }
            }
            
            return null;
        }
    });

    // Register diagnostics
    const diagnosticCollection = vscode.languages.createDiagnosticCollection('banglacode');
    
    const updateDiagnostics = (document) => {
        if (document.languageId !== 'banglacode') return;
        
        const diagnostics = [];
        const text = document.getText();
        
        // Find all declared variables
        const declaredVars = new Set();
        const varDeclRegex = /dhoro\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*=/g;
        let match;
        while ((match = varDeclRegex.exec(text)) !== null) {
            declaredVars.add(match[1]);
        }
        
        // Find all function definitions with their parameter counts
        const funcDefs = new Map();
        const funcDefRegex = /kaj\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)/g;
        while ((match = funcDefRegex.exec(text)) !== null) {
            const name = match[1];
            const params = match[2].trim();
            const paramCount = params ? params.split(',').length : 0;
            funcDefs.set(name, { params, paramCount });
        }
        
        // Find all function calls and check argument counts
        const funcCallRegex = /([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)/g;
        while ((match = funcCallRegex.exec(text)) !== null) {
            const funcName = match[1];
            const argsStr = match[2].trim();
            
            // Skip if it's a function definition, built-in, or keyword
            if (funcName === 'kaj' || funcName === 'jodi' || funcName === 'jotokkhon' || 
                funcName === 'ghuriye' || funcName === 'shuru' ||
                builtinFunctions.some(f => f.label === funcName)) {
                continue;
            }
            
            if (funcDefs.has(funcName)) {
                const def = funcDefs.get(funcName);
                // Count arguments (simple approach - count commas + 1, unless empty)
                const argCount = argsStr ? argsStr.split(',').length : 0;
                
                if (argCount !== def.paramCount) {
                    const pos = document.positionAt(match.index);
                    const range = new vscode.Range(pos, pos.translate(0, funcName.length));
                    diagnostics.push(new vscode.Diagnostic(
                        range,
                        `Function '${funcName}' expects ${def.paramCount} argument(s) but got ${argCount}`,
                        vscode.DiagnosticSeverity.Error
                    ));
                }
            }
        }
        
        // Check for aliased module function calls (e.g., math.add(5))
        // First, find all import aliases and their modules
        const importAliases = new Map();
        const aliasImportRegex = /ano\s+"([^"]+)"\s+hisabe\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*;?/g;
        while ((match = aliasImportRegex.exec(text)) !== null) {
            importAliases.set(match[2], match[1]);
        }
        
        // Find alias.function() calls
        const aliasCallRegex = /([a-zA-Z_][a-zA-Z0-9_]*)\.([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)/g;
        while ((match = aliasCallRegex.exec(text)) !== null) {
            const aliasName = match[1];
            const funcName = match[2];
            const argsStr = match[3].trim();
            
            if (importAliases.has(aliasName)) {
                const modulePath = importAliases.get(aliasName);
                try {
                    const currentDir = path.dirname(document.uri.fsPath);
                    const fullPath = path.resolve(currentDir, modulePath);
                    if (fs.existsSync(fullPath)) {
                        const moduleText = fs.readFileSync(fullPath, 'utf8');
                        
                        // Find function definition in module
                        const funcInModuleRegex = new RegExp(`pathao\\s+kaj\\s+${funcName}\\s*\\(([^)]*)\\)`);
                        const funcInModule = funcInModuleRegex.exec(moduleText);
                        
                        if (funcInModule) {
                            const params = funcInModule[1].trim();
                            const expectedCount = params ? params.split(',').length : 0;
                            const actualCount = argsStr ? argsStr.split(',').length : 0;
                            
                            if (actualCount !== expectedCount) {
                                const pos = document.positionAt(match.index);
                                const range = new vscode.Range(pos, pos.translate(0, aliasName.length + 1 + funcName.length));
                                diagnostics.push(new vscode.Diagnostic(
                                    range,
                                    `Function '${aliasName}.${funcName}' expects ${expectedCount} argument(s) but got ${actualCount}`,
                                    vscode.DiagnosticSeverity.Error
                                ));
                            }
                        }
                    }
                } catch (err) {
                    // Ignore errors
                }
            }
        }
        
        diagnosticCollection.set(document.uri, diagnostics);
    };
    
    // Update diagnostics on document change
    vscode.workspace.onDidChangeTextDocument(event => {
        updateDiagnostics(event.document);
    });
    
    // Update diagnostics on document open
    vscode.workspace.onDidOpenTextDocument(document => {
        updateDiagnostics(document);
    });
    
    // Initial diagnostics for open documents
    vscode.workspace.textDocuments.forEach(document => {
        updateDiagnostics(document);
    });

    // Register signature help provider
    const signatureProvider = vscode.languages.registerSignatureHelpProvider(
        'banglacode',
        {
            provideSignatureHelp(document, position, token, context) {
                const line = document.lineAt(position.line).text;
                const beforeCursor = line.substring(0, position.character);
                
                // Find function name
                const match = beforeCursor.match(/([a-zA-Z_][a-zA-Z0-9_]*)\s*\([^)]*$/);
                if (!match) return null;
                
                const funcName = match[1];
                const fn = builtinFunctions.find(f => f.label === funcName);
                if (!fn) return null;
                
                const sigHelp = new vscode.SignatureHelp();
                const sig = new vscode.SignatureInformation(fn.insertText.replace(/\$\{\d+:?/g, '').replace(/\}/g, ''), fn.documentation);
                sigHelp.signatures = [sig];
                sigHelp.activeSignature = 0;
                
                return sigHelp;
            }
        },
        '(', ','
    );

    context.subscriptions.push(completionProvider, importPathProvider, hoverProvider, signatureProvider, diagnosticCollection);
}

function deactivate() {}

module.exports = { activate, deactivate };
