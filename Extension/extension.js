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
    { label: 'dhoro', kind: vscode.CompletionItemKind.Keyword, detail: 'চলক ঘোষণা (ধরো) - Variable', insertText: 'dhoro ${1:name} = ${2:value};' },
    { label: 'sthir', kind: vscode.CompletionItemKind.Keyword, detail: 'স্থির মান (স্থির) - Constant', insertText: 'sthir ${1:NAME} = ${2:value};' },
    { label: 'bishwo', kind: vscode.CompletionItemKind.Keyword, detail: 'বৈশ্বিক চলক (বিশ্ব) - Global', insertText: 'bishwo ${1:name} = ${2:value};' },
    { label: 'jodi', kind: vscode.CompletionItemKind.Keyword, detail: 'শর্ত যদি (যদি) - If', insertText: 'jodi (${1:condition}) {\n\t${2}\n}' },
    { label: 'nahole', kind: vscode.CompletionItemKind.Keyword, detail: 'নাহলে (নাহলে) - Else', insertText: 'nahole {\n\t${1}\n}' },
    { label: 'jotokkhon', kind: vscode.CompletionItemKind.Keyword, detail: 'যতক্ষণ লুপ (যতক্ষণ) - While', insertText: 'jotokkhon (${1:condition}) {\n\t${2}\n}' },
    { label: 'ghuriye', kind: vscode.CompletionItemKind.Keyword, detail: 'ঘুরিয়ে লুপ (ঘুরিয়ে) - For', insertText: 'ghuriye (dhoro ${1:i} = 0; ${1:i} < ${2:10}; ${1:i} = ${1:i} + 1) {\n\t${3}\n}' },
    { label: 'kaj', kind: vscode.CompletionItemKind.Keyword, detail: 'ফাংশন (কাজ) - Function', insertText: 'kaj ${1:name}(${2:params}) {\n\t${3}\n}' },
    { label: 'ferao', kind: vscode.CompletionItemKind.Keyword, detail: 'ফেরাও (ফেরাও) - Return', insertText: 'ferao ${1:value};' },
    { label: 'sreni', kind: vscode.CompletionItemKind.Keyword, detail: 'শ্রেণী (শ্রেণী) - Class', insertText: 'sreni ${1:Name} {\n\tshuru(${2:params}) {\n\t\t${3}\n\t}\n}' },
    { label: 'shuru', kind: vscode.CompletionItemKind.Keyword, detail: 'কন্সট্রাক্টর (শুরু) - Constructor', insertText: 'shuru(${1:params}) {\n\tei.${2:property} = ${2:property};\n}' },
    { label: 'notun', kind: vscode.CompletionItemKind.Keyword, detail: 'নতুন অবজেক্ট (নতুন) - New', insertText: 'notun ${1:ClassName}(${2:args})' },
    { label: 'sotti', kind: vscode.CompletionItemKind.Constant, detail: 'সত্যি (সত্যি) - True', insertText: 'sotti' },
    { label: 'mittha', kind: vscode.CompletionItemKind.Constant, detail: 'মিথ্যা (মিথ্যা) - False', insertText: 'mittha' },
    { label: 'khali', kind: vscode.CompletionItemKind.Constant, detail: 'খালি (খালি) - Null', insertText: 'khali' },
    { label: 'ebong', kind: vscode.CompletionItemKind.Operator, detail: 'এবং (এবং) - AND', insertText: 'ebong' },
    { label: 'ba', kind: vscode.CompletionItemKind.Operator, detail: 'বা (বা) - OR', insertText: 'ba' },
    { label: 'na', kind: vscode.CompletionItemKind.Operator, detail: 'না (না) - NOT', insertText: 'na' },
    { label: 'thamo', kind: vscode.CompletionItemKind.Keyword, detail: 'থামো (থামো) - Break', insertText: 'thamo;' },
    { label: 'chharo', kind: vscode.CompletionItemKind.Keyword, detail: 'ছাড়ো (ছাড়ো) - Continue', insertText: 'chharo;' },
    { label: 'ei', kind: vscode.CompletionItemKind.Keyword, detail: 'এই (এই) - This', insertText: 'ei.' },
    { label: 'ano', kind: vscode.CompletionItemKind.Keyword, detail: 'আনো মডিউল (আনো) - Import', insertText: 'ano "${1:module.bang}";' },
    { label: 'hisabe', kind: vscode.CompletionItemKind.Keyword, detail: 'হিসাবে (হিসাবে) - As', insertText: 'hisabe ${1:alias}' },
    { label: 'pathao', kind: vscode.CompletionItemKind.Keyword, detail: 'পাঠাও (পাঠাও) - Export', insertText: 'pathao ' },
    { label: 'chesta', kind: vscode.CompletionItemKind.Keyword, detail: 'চেষ্টা করো (চেষ্টা) - Try', insertText: 'chesta {\n\t${1}\n} dhoro_bhul (${2:err}) {\n\t${3}\n}' },
    { label: 'dhoro_bhul', kind: vscode.CompletionItemKind.Keyword, detail: 'ভুল ধরো (ধরো ভুল) - Catch', insertText: 'dhoro_bhul (${1:err}) {\n\t${2}\n}' },
    { label: 'shesh', kind: vscode.CompletionItemKind.Keyword, detail: 'শেষে (শেষ) - Finally', insertText: 'shesh {\n\t${1}\n}' },
    { label: 'felo', kind: vscode.CompletionItemKind.Keyword, detail: 'ত্রুটি ফেলো (ফেলো) - Throw', insertText: 'felo "${1:error message}";' },
];

// Built-in functions
const builtinFunctions = [
    // আউটপুট
    { label: 'dekho', kind: vscode.CompletionItemKind.Function, detail: 'দেখো - মান প্রিন্ট করো', insertText: 'dekho(${1:value});', documentation: 'কনসোলে মান প্রিন্ট করে' },

    // টাইপ ফাংশন
    { label: 'dhoron', kind: vscode.CompletionItemKind.Function, detail: 'ধরন - ডেটা টাইপ জানো', insertText: 'dhoron(${1:value})', documentation: 'মানের ধরন রিটার্ন করে' },
    { label: 'lipi', kind: vscode.CompletionItemKind.Function, detail: 'লিপি - স্ট্রিং এ রূপান্তর', insertText: 'lipi(${1:value})', documentation: 'মানকে স্ট্রিং এ রূপান্তর করে' },
    { label: 'sonkha', kind: vscode.CompletionItemKind.Function, detail: 'সংখ্যা - নাম্বার এ রূপান্তর', insertText: 'sonkha(${1:value})', documentation: 'মানকে সংখ্যায় রূপান্তর করে' },
    { label: 'dorghyo', kind: vscode.CompletionItemKind.Function, detail: 'দৈর্ঘ্য - লেন্থ নাও', insertText: 'dorghyo(${1:value})', documentation: 'স্ট্রিং বা অ্যারের দৈর্ঘ্য রিটার্ন করে' },

    // স্ট্রিং ফাংশন
    { label: 'boroHater', kind: vscode.CompletionItemKind.Function, detail: 'বড় হাতের - আপারকেস', insertText: 'boroHater(${1:str})', documentation: 'স্ট্রিংকে বড় হাতের অক্ষরে রূপান্তর করে' },
    { label: 'chotoHater', kind: vscode.CompletionItemKind.Function, detail: 'ছোট হাতের - লোয়ারকেস', insertText: 'chotoHater(${1:str})', documentation: 'স্ট্রিংকে ছোট হাতের অক্ষরে রূপান্তর করে' },
    { label: 'chhanto', kind: vscode.CompletionItemKind.Function, detail: 'ছাঁটো - ট্রিম করো', insertText: 'chhanto(${1:str})', documentation: 'স্ট্রিং থেকে ফাঁকা জায়গা সরায়' },
    { label: 'bhag', kind: vscode.CompletionItemKind.Function, detail: 'ভাগ - স্প্লিট করো', insertText: 'bhag(${1:str}, ${2:separator})', documentation: 'স্ট্রিংকে অ্যারেতে ভাগ করে' },
    { label: 'joro', kind: vscode.CompletionItemKind.Function, detail: 'জোড়ো - জয়েন করো', insertText: 'joro(${1:arr}, ${2:separator})', documentation: 'অ্যারেকে স্ট্রিং এ জোড়ে' },
    { label: 'khojo', kind: vscode.CompletionItemKind.Function, detail: 'খোঁজো - ইনডেক্স খুঁজো', insertText: 'khojo(${1:str}, ${2:substr})', documentation: 'সাবস্ট্রিংয়ের ইনডেক্স খুঁজে দেয় (-১ যদি না পাওয়া যায়)' },
    { label: 'angsho', kind: vscode.CompletionItemKind.Function, detail: 'অংশ - সাবস্ট্রিং নাও', insertText: 'angsho(${1:str}, ${2:start}, ${3:end})', documentation: 'শুরু থেকে শেষ পর্যন্ত সাবস্ট্রিং বের করে' },
    { label: 'bodlo', kind: vscode.CompletionItemKind.Function, detail: 'বদলো - রিপ্লেস করো', insertText: 'bodlo(${1:str}, ${2:old}, ${3:new})', documentation: 'সব জায়গায় পুরনো মানকে নতুন দিয়ে বদলায়' },

    // অ্যারে ফাংশন
    { label: 'dhokao', kind: vscode.CompletionItemKind.Function, detail: 'ঢোকাও - পুশ করো', insertText: 'dhokao(${1:arr}, ${2:value})', documentation: 'অ্যারেতে এলিমেন্ট যোগ করে' },
    { label: 'berKoro', kind: vscode.CompletionItemKind.Function, detail: 'বের করো - পপ করো', insertText: 'berKoro(${1:arr})', documentation: 'শেষ এলিমেন্ট সরিয়ে রিটার্ন করে' },
    { label: 'kato', kind: vscode.CompletionItemKind.Function, detail: 'কাটো - স্লাইস করো', insertText: 'kato(${1:arr}, ${2:start}, ${3:end})', documentation: 'অ্যারের একটি অংশ বের করে' },
    { label: 'ulto', kind: vscode.CompletionItemKind.Function, detail: 'উল্টো - রিভার্স করো', insertText: 'ulto(${1:arr})', documentation: 'অ্যারে উল্টে দেয় (নতুন অ্যারে রিটার্ন করে)' },
    { label: 'saja', kind: vscode.CompletionItemKind.Function, detail: 'সাজা - সর্ট করো', insertText: 'saja(${1:arr})', documentation: 'অ্যারে সাজায় (নতুন অ্যারে রিটার্ন করে)' },
    { label: 'ache', kind: vscode.CompletionItemKind.Function, detail: 'আছে - খুঁজে দেখো', insertText: 'ache(${1:arr}, ${2:value})', documentation: 'অ্যারেতে মান আছে কিনা চেক করে' },
    { label: 'chabi', kind: vscode.CompletionItemKind.Function, detail: 'চাবি - কী গুলো নাও', insertText: 'chabi(${1:map})', documentation: 'ম্যাপের সব কী এর অ্যারে রিটার্ন করে' },

    // গাণিতিক ফাংশন
    { label: 'borgomul', kind: vscode.CompletionItemKind.Function, detail: 'বর্গমূল - স্কয়ার রুট', insertText: 'borgomul(${1:x})', documentation: 'বর্গমূল রিটার্ন করে' },
    { label: 'ghat', kind: vscode.CompletionItemKind.Function, detail: 'ঘাত - পাওয়ার', insertText: 'ghat(${1:base}, ${2:exp})', documentation: 'বেস কে এক্সপোনেন্ট এ উন্নীত করে' },
    { label: 'niche', kind: vscode.CompletionItemKind.Function, detail: 'নিচে - ফ্লোর', insertText: 'niche(${1:x})', documentation: 'নিচের দিকে রাউন্ড করে' },
    { label: 'upore', kind: vscode.CompletionItemKind.Function, detail: 'উপরে - সিলিং', insertText: 'upore(${1:x})', documentation: 'উপরের দিকে রাউন্ড করে' },
    { label: 'kache', kind: vscode.CompletionItemKind.Function, detail: 'কাছে - রাউন্ড', insertText: 'kache(${1:x})', documentation: 'কাছের পূর্ণসংখ্যায় রাউন্ড করে' },
    { label: 'niratek', kind: vscode.CompletionItemKind.Function, detail: 'নিরপেক্ষ - এবসোলিউট', insertText: 'niratek(${1:x})', documentation: 'নিরপেক্ষ মান রিটার্ন করে' },
    { label: 'choto', kind: vscode.CompletionItemKind.Function, detail: 'ছোট - মিনিমাম', insertText: 'choto(${1:a}, ${2:b})', documentation: 'ছোট মান রিটার্ন করে' },
    { label: 'boro', kind: vscode.CompletionItemKind.Function, detail: 'বড় - ম্যাক্সিমাম', insertText: 'boro(${1:a}, ${2:b})', documentation: 'বড় মান রিটার্ন করে' },
    { label: 'lotto', kind: vscode.CompletionItemKind.Function, detail: 'লটো - র্যান্ডম', insertText: 'lotto()', documentation: '০ থেকে ১ এর মধ্যে র্যান্ডম নাম্বার রিটার্ন করে' },

    // ইউটিলিটি ফাংশন
    { label: 'somoy', kind: vscode.CompletionItemKind.Function, detail: 'সময় - বর্তমান সময়', insertText: 'somoy()', documentation: 'বর্তমান টাইমস্ট্যাম্প মিলিসেকেন্ডে রিটার্ন করে' },
    { label: 'ghum', kind: vscode.CompletionItemKind.Function, detail: 'ঘুম - স্লিপ করো', insertText: 'ghum(${1:milliseconds})', documentation: 'নির্দিষ্ট মিলিসেকেন্ড এর জন্য থামিয়ে রাখে' },
    { label: 'nao', kind: vscode.CompletionItemKind.Function, detail: 'নাও - ইনপুট নাও', insertText: 'nao(${1:"মান লেখো: "})', documentation: 'কনসোল থেকে ইউজার ইনপুট পড়ে' },
    { label: 'bondho', kind: vscode.CompletionItemKind.Function, detail: 'বন্ধ - এক্সিট করো', insertText: 'bondho(${1:0})', documentation: 'প্রোগ্রাম বন্ধ করে' },

    // ফাইল ফাংশন
    { label: 'poro', kind: vscode.CompletionItemKind.Function, detail: 'পড়ো - ফাইল পড়ো', insertText: 'poro(${1:"filename"})', documentation: 'ফাইল এর কন্টেন্ট স্ট্রিং হিসেবে পড়ে' },
    { label: 'lekho', kind: vscode.CompletionItemKind.Function, detail: 'লেখো - ফাইল লেখো', insertText: 'lekho(${1:"filename"}, ${2:content})', documentation: 'ফাইলে কন্টেন্ট লেখে' },

    // HTTP ফাংশন
    { label: 'server_chalu', kind: vscode.CompletionItemKind.Function, detail: 'সার্ভার চালু - HTTP সার্ভার', insertText: 'server_chalu(${1:3000}, ${2:handler})', documentation: 'নির্দিষ্ট পোর্টে HTTP সার্ভার চালু করে' },
    { label: 'anun', kind: vscode.CompletionItemKind.Function, detail: 'আনুন - HTTP GET রিকোয়েস্ট', insertText: 'anun(${1:"url"})', documentation: 'HTTP GET রিকোয়েস্ট করে' },
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
 * Extract return statement from a function body
 * @param {string} text - Full document text
 * @param {string} funcName - Function name
 * @returns {string|null} - Return expression or null
 */
function extractReturnValue(text, funcName) {
    // Find the function definition and extract its body
    const funcDefRegex = new RegExp(`(?:pathao\\s+)?kaj\\s+${funcName}\\s*\\([^)]*\\)\\s*\\{`, 'g');
    const funcMatch = funcDefRegex.exec(text);

    if (!funcMatch) return null;

    const startIndex = funcMatch.index + funcMatch[0].length - 1; // Position of opening brace
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

    const funcBody = text.substring(startIndex + 1, endIndex);

    // Look for ferao (return) statement
    const returnMatch = funcBody.match(/ferao\s+([^;]+);?/);
    if (returnMatch) {
        return returnMatch[1].trim();
    }

    return null;
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

    // Find the map declaration: dhoro/sthir/bishwo varName = { ... }
    // Use a simple brace-matching approach
    const startRegex = new RegExp(`(dhoro|sthir|bishwo)\\s+${varName}\\s*=\\s*\\{`);
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

/**
 * Extract variable declaration info including value and inferred type
 * @param {string} text - Full document text
 * @param {string} varName - Variable name to find
 * @returns {object|null} - { declType, value, inferredType, displayValue }
 */
function extractVariableInfo(text, varName) {
    // Match: (dhoro|sthir|bishwo) varName = value;
    const declRegex = new RegExp(`(dhoro|sthir|bishwo)\\s+${varName}\\s*=\\s*([^;]+);?`, 'm');
    const match = declRegex.exec(text);
    if (!match) return null;

    const declType = match[1];
    let rawValue = match[2].trim();
    let inferredType = 'unknown';
    let displayValue = rawValue;

    // Infer type from value
    if (rawValue.startsWith('"') || rawValue.startsWith("'")) {
        inferredType = 'string';
        // Extract string content for display (remove quotes)
        const strMatch = rawValue.match(/^["'](.*)["']$/);
        if (strMatch) {
            displayValue = `"${strMatch[1]}"`;
        }
    } else if (rawValue === 'sotti') {
        inferredType = 'boolean';
        displayValue = 'sotti (true)';
    } else if (rawValue === 'mittha') {
        inferredType = 'boolean';
        displayValue = 'mittha (false)';
    } else if (rawValue === 'khali') {
        inferredType = 'null';
        displayValue = 'khali (null)';
    } else if (rawValue.startsWith('[')) {
        inferredType = 'array';
        // Try to show array length or truncated preview
        const arrayContent = rawValue.match(/^\[([\s\S]*)\]$/);
        if (arrayContent) {
            const items = arrayContent[1].split(',').filter(s => s.trim());
            displayValue = `[...] (${items.length} items)`;
        }
    } else if (rawValue.startsWith('{')) {
        inferredType = 'map';
        displayValue = '{...}';
    } else if (/^-?\d+(\.\d+)?$/.test(rawValue)) {
        inferredType = rawValue.includes('.') ? 'number (float)' : 'number (int)';
        displayValue = rawValue;
    } else if (rawValue.startsWith('kaj')) {
        inferredType = 'function';
        displayValue = 'kaj () {...}';
    } else if (rawValue.startsWith('notun ')) {
        // Class instance: notun ClassName(...)
        const classMatch = rawValue.match(/^notun\s+([a-zA-Z_][a-zA-Z0-9_]*)/);
        if (classMatch) {
            inferredType = classMatch[1];
            displayValue = rawValue;
        }
    } else if (/^[a-zA-Z_][a-zA-Z0-9_]*\s*\(/.test(rawValue)) {
        // Function call result
        const funcMatch = rawValue.match(/^([a-zA-Z_][a-zA-Z0-9_]*)\s*\(/);
        if (funcMatch) {
            inferredType = `result of ${funcMatch[1]}()`;
        }
    }

    return {
        declType,
        value: rawValue,
        inferredType,
        displayValue
    };
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
                const varRegex = /(dhoro|sthir|bishwo)\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*=/g;
                let match;
                const seenVars = new Set();
                while ((match = varRegex.exec(text)) !== null) {
                    const declType = match[1];
                    const varName = match[2];
                    if (!seenVars.has(varName)) {
                        seenVars.add(varName);
                        const kind = declType === 'sthir' ? vscode.CompletionItemKind.Constant : vscode.CompletionItemKind.Variable;
                        const item = new vscode.CompletionItem(varName, kind);
                        item.detail = declType === 'sthir' ? 'স্থির মান - Constant' : (declType === 'bishwo' ? 'বৈশ্বিক চলক - Global' : 'চলক - Variable');
                        completions.push(item);
                    }
                }
                
                // Extract functions from current document
                const funcRegex = /kaj\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)/g;
                const seenFuncs = new Set();
                while ((match = funcRegex.exec(text)) !== null) {
                    const funcName = match[1];
                    const params = match[2].trim();
                    if (!seenFuncs.has(funcName) && funcName !== 'shuru') {
                        seenFuncs.add(funcName);
                        const returnValue = extractReturnValue(text, funcName);
                        const item = new vscode.CompletionItem(funcName, vscode.CompletionItemKind.Function);
                        const returnInfo = returnValue ? ` → ${returnValue}` : '';
                        item.detail = `kaj ${funcName}(${params})${returnInfo}`;
                        item.insertText = new vscode.SnippetString(funcName + '(${1})');
                        if (returnValue) {
                            item.documentation = new vscode.MarkdownString(`**রিটার্ন করে:** \`${returnValue}\``);
                        }
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
                        item.detail = 'ইউজার শ্রেণী - User Class';
                        completions.push(item);
                    }
                }

                // Extract functions from directly imported modules (ano "..." without alias)
                const directImportRegex = /ano\s+["']([^"']+)["']\s*;/g;
                let directImportMatch;
                const seenImportedFuncs = new Set();

                while ((directImportMatch = directImportRegex.exec(text)) !== null) {
                    const modulePath = directImportMatch[1];

                    // Skip if this import has an alias (those are handled by typing alias.)
                    const hasAlias = new RegExp(`ano\\s+["']${modulePath.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')}["']\\s+hisabe`, 'm').test(text);
                    if (hasAlias) {
                        continue;
                    }

                    // Only process .bang modules
                    if (!modulePath.endsWith('.bang')) {
                        continue;
                    }

                    try {
                        const currentDir = path.dirname(document.uri.fsPath);
                        const fullPath = path.resolve(currentDir, modulePath);

                        if (!fs.existsSync(fullPath)) {
                            continue;
                        }

                        const moduleText = fs.readFileSync(fullPath, 'utf8');

                        // Find all exported functions (with 'pathao' keyword)
                        const exportedFuncRegex = /pathao\s+kaj\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)/g;
                        let exportMatch;

                        while ((exportMatch = exportedFuncRegex.exec(moduleText)) !== null) {
                            const funcName = exportMatch[1];
                            const params = exportMatch[2].trim();

                            if (funcName === 'shuru' || seenImportedFuncs.has(funcName)) {
                                continue;
                            }
                            seenImportedFuncs.add(funcName);

                            const docComment = extractDocComment(moduleText, funcName, 'kaj');
                            const returnValue = extractReturnValue(moduleText, funcName);
                            const paramList = params ? params.split(',').map(p => p.trim()) : [];

                            const item = new vscode.CompletionItem(funcName, vscode.CompletionItemKind.Function);
                            const returnInfo = returnValue ? ` → ${returnValue}` : '';
                            item.detail = `kaj ${funcName}(${params})${returnInfo} - ইম্পোর্ট করা`;
                            item.insertText = new vscode.SnippetString(funcName + '(${1})');
                            item.sortText = `1_${funcName}`; // Lower priority than local functions

                            // Build documentation
                            const docMd = new vscode.MarkdownString();
                            docMd.appendMarkdown(`**kaj ${funcName}(${params})**\n\n`);
                            docMd.appendMarkdown(`📦 ইম্পোর্ট করা: \`${modulePath}\`\n\n`);

                            if (docComment) {
                                docMd.appendMarkdown('---\n\n');
                                docMd.appendMarkdown(docComment + '\n\n');
                                docMd.appendMarkdown('---\n\n');
                            }

                            docMd.appendMarkdown(`**প্যারামিটার:** ${paramList.length === 0 ? 'নেই' : paramList.join(', ')}\n\n`);

                            if (returnValue) {
                                docMd.appendMarkdown(`**রিটার্ন করে:** \`${returnValue}\`\n\n`);
                            }

                            docMd.appendMarkdown(`**কল করো:** \`${funcName}(${paramList.map((_, i) => `arg${i+1}`).join(', ')})\``);

                            item.documentation = docMd;
                            completions.push(item);
                        }
                    } catch (err) {
                        console.error('[BanglaCode] Error reading module for completion:', err);
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
                            item.detail = 'ডিরেক্টরি';
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
                            item.detail = 'BanglaCode মডিউল';
                            item.insertText = entry.name;
                            completions.push(item);
                        } else if (entry.name.endsWith('.json')) {
                            const item = new vscode.CompletionItem(
                                entry.name,
                                vscode.CompletionItemKind.File
                            );
                            item.detail = 'JSON ফাইল';
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

    // Register map/JSON property completion provider (also handles module imports)
    const mapPropertyProvider = vscode.languages.registerCompletionItemProvider(
        'banglacode',
        {
            provideCompletionItems(document, position, token, context) {
                const line = document.lineAt(position.line).text;
                const linePrefix = line.substring(0, position.character);

                // Check if we're after a dot (e.g., "json." or "as.")
                const dotMatch = linePrefix.match(/([a-zA-Z_][a-zA-Z0-9_]*)\.$/);
                if (!dotMatch) {
                    return undefined;
                }

                const varName = dotMatch[1];
                const text = document.getText();
                const completions = [];

                // PRIORITY 1: Check if this is an import alias (ano "..." hisabe varName)
                const aliasImportRegex = new RegExp(`ano\\s+["']([^"']+)["']\\s+hisabe\\s+${varName}\\s*;?`, 'm');
                const aliasImport = aliasImportRegex.exec(text);

                if (aliasImport) {
                    const modulePath = aliasImport[1];
                    try {
                        const currentDir = path.dirname(document.uri.fsPath);
                        const fullPath = path.resolve(currentDir, modulePath);

                        if (!fs.existsSync(fullPath)) {
                            console.log('[BanglaCode] Module file not found:', fullPath);
                            return completions;
                        }

                        // Check if it's a JSON file
                        if (modulePath.endsWith('.json')) {
                            const jsonText = fs.readFileSync(fullPath, 'utf8');
                            try {
                                const jsonData = JSON.parse(jsonText);
                                const keys = extractJSONKeys(jsonData);
                                for (const key of keys) {
                                    const item = new vscode.CompletionItem(key.name, vscode.CompletionItemKind.Property);
                                    item.detail = `${key.type} প্রপার্টি (JSON)`;
                                    item.sortText = `0_${key.name}`; // Prioritize in completion list
                                    if (key.nested) {
                                        item.documentation = new vscode.MarkdownString('**নেস্টেড অবজেক্ট:** ' + key.nested.join(', '));
                                    }
                                    completions.push(item);
                                }
                            } catch (err) {
                                console.error('[BanglaCode] Error parsing JSON:', err);
                            }
                        } else if (modulePath.endsWith('.bang')) {
                            // It's a .bang module - get ONLY exported functions (with 'pathao' keyword)
                            const moduleText = fs.readFileSync(fullPath, 'utf8');

                            // Find only exported functions (with 'pathao' keyword)
                            const funcRegex = /pathao\s+kaj\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)/g;
                            let funcMatch;
                            const seenFuncs = new Set();

                            while ((funcMatch = funcRegex.exec(moduleText)) !== null) {
                                const funcName = funcMatch[1];
                                if (funcName === 'shuru' || seenFuncs.has(funcName)) continue; // Skip constructor and duplicates
                                seenFuncs.add(funcName);

                                const params = funcMatch[2].trim();
                                const paramList = params ? params.split(',').map(p => p.trim()) : [];
                                const docComment = extractDocComment(moduleText, funcName, 'kaj');
                                const returnValue = extractReturnValue(moduleText, funcName);

                                const item = new vscode.CompletionItem(funcName, vscode.CompletionItemKind.Method);
                                const returnInfo = returnValue ? ` → ${returnValue}` : '';
                                item.detail = `kaj ${funcName}(${params})${returnInfo}`;
                                item.insertText = new vscode.SnippetString(funcName + '(${1})');
                                item.sortText = `0_${funcName}`; // Prioritize in completion list

                                // Build documentation
                                const docMd = new vscode.MarkdownString();
                                docMd.appendMarkdown(`**${varName}.${funcName}(${params})**\n\n`);
                                docMd.appendMarkdown(`📦 ইম্পোর্ট করা: \`${modulePath}\`\n\n`);

                                if (docComment) {
                                    docMd.appendMarkdown('---\n\n');
                                    docMd.appendMarkdown(docComment + '\n\n');
                                    docMd.appendMarkdown('---\n\n');
                                }

                                docMd.appendMarkdown(`**প্যারামিটার:** ${paramList.length === 0 ? 'নেই' : paramList.join(', ')}\n\n`);

                                if (returnValue) {
                                    docMd.appendMarkdown(`**রিটার্ন করে:** \`${returnValue}\`\n\n`);
                                }

                                docMd.appendMarkdown(`**কল করো:** \`${varName}.${funcName}(${paramList.map((_, i) => `arg${i+1}`).join(', ')})\``);

                                item.documentation = docMd;
                                completions.push(item);
                            }

                            console.log(`[BanglaCode] Found ${completions.length} functions in ${modulePath} for alias '${varName}'`);
                        }
                    } catch (err) {
                        console.error('[BanglaCode] Error reading module:', err);
                    }

                    return completions;
                }

                // PRIORITY 2: Check if it's a map/JSON variable - extract its keys
                const mapKeys = extractMapKeys(text, varName);
                if (mapKeys.length > 0) {
                    for (const key of mapKeys) {
                        const item = new vscode.CompletionItem(key.name, vscode.CompletionItemKind.Property);
                        item.detail = key.type ? `প্রপার্টি (${key.type})` : 'প্রপার্টি';
                        item.sortText = `1_${key.name}`;
                        if (key.nested) {
                            item.documentation = new vscode.MarkdownString('**নেস্টেড অবজেক্ট প্রপার্টি:** ' + key.nested.join(', '));
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
                    'BanglaCode কীওয়ার্ড'
                ]);
            }

            // Check built-in functions
            const fn = builtinFunctions.find(f => f.label === word);
            if (fn) {
                return new vscode.Hover([
                    `**${fn.label}** - ${fn.detail}`,
                    fn.documentation || '',
                    'BanglaCode বিল্ট-ইন ফাংশন'
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

                // Extract return value
                const returnValue = extractReturnValue(text, word);

                const md = new vscode.MarkdownString();
                md.appendMarkdown(`**kaj ${word}(${params})**\n\n`);

                if (docComment) {
                    md.appendMarkdown('---\n\n');
                    md.appendMarkdown(docComment + '\n\n');
                    md.appendMarkdown('---\n\n');
                }

                md.appendMarkdown(`**প্যারামিটার:** ${paramCount === 0 ? 'নেই' : paramList.join(', ')}\n\n`);

                if (returnValue) {
                    md.appendMarkdown(`**রিটার্ন করে:** \`${returnValue}\`\n\n`);
                } else {
                    md.appendMarkdown(`**রিটার্ন করে:** কিছু না (void)\n\n`);
                }

                md.appendMarkdown(`**কল করো:** \`${word}(${paramList.map((_, i) => `arg${i+1}`).join(', ')})\``);

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
                            md.appendMarkdown(`**${word}** - ইম্পোর্ট করা JSON\n\n`);
                            md.appendMarkdown(`📦 ফাইল: \`${modulePath}\`\n\n`);
                            md.appendMarkdown('---\n\n');
                            md.appendMarkdown('**প্রপার্টি:**\n\n');
                            for (const key of keys.slice(0, 10)) {
                                const value = jsonData[key];
                                const type = Array.isArray(value) ? 'array' : typeof value;
                                md.appendMarkdown(`• \`${key}\`: ${type}\n\n`);
                            }
                            if (keys.length > 10) {
                                md.appendMarkdown(`... এবং আরও ${keys.length - 10} টি\n\n`);
                            }
                            md.appendMarkdown('---\n\n');
                            md.appendMarkdown(`**ব্যবহার করো:** \`${word}.propertyName\``);
                            
                            return new vscode.Hover(md);
                        }
                    } catch (err) {
                        // Fallback to basic hover
                    }
                }
                
                // Regular .bang module
                return new vscode.Hover([
                    `**${word}** - ইম্পোর্ট করা মডিউলের নাম`,
                    `মডিউল: \`${modulePath}\``,
                    `ব্যবহার করো: ${word}.functionName(args)`
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
                
                md.appendMarkdown(constructorParams ? `**কন্সট্রাক্টর:** shuru(${constructorParams})\n\n` : 'কোনো কন্সট্রাক্টর নেই\n\n');
                md.appendMarkdown(`**তৈরি করো:** \`notun ${word}(${constructorParams})\``);
                
                return new vscode.Hover(md);
            }
            
            // Check variables (dhoro, sthir, bishwo)
            const varInfo = extractVariableInfo(text, word);
            if (varInfo) {
                const { declType, inferredType, displayValue } = varInfo;
                const banglaLabel = declType === 'sthir' ? 'স্থির' : (declType === 'bishwo' ? 'বিশ্ব' : 'ধরো');

                // Check if it's a map/JSON variable
                const mapKeys = extractMapKeys(text, word);
                if (mapKeys.length > 0) {
                    const md = new vscode.MarkdownString();
                    md.isTrusted = true;
                    md.appendMarkdown(`**${word}** - ম্যাপ/JSON অবজেক্ট\n\n`);
                    md.appendMarkdown(`**ঘোষণা করা হয়েছে:** ${declType} (${banglaLabel})\n\n`);
                    md.appendMarkdown('---\n\n');
                    md.appendMarkdown('**প্রপার্টি:**\n\n');
                    for (const key of mapKeys) {
                        if (key.nested) {
                            md.appendMarkdown(`• \`${key.name}\`: object { ${key.nested.join(', ')} }\n\n`);
                        } else {
                            md.appendMarkdown(`• \`${key.name}\`: ${key.type}\n\n`);
                        }
                    }
                    md.appendMarkdown('---\n\n');
                    md.appendMarkdown(`**ব্যবহার করো:** \`${word}.propertyName\``);
                    return new vscode.Hover(md);
                }

                // Standard variable hover with type and value
                const md = new vscode.MarkdownString();
                md.isTrusted = true;
                md.appendMarkdown(`**${word}** - ${inferredType}\n\n`);
                md.appendMarkdown(`**ঘোষণা করা হয়েছে:** ${declType} (${banglaLabel})`);
                if (declType === 'sthir') {
                    md.appendMarkdown(` - স্থির মান (অপরিবর্তনীয়)`);
                } else if (declType === 'bishwo') {
                    md.appendMarkdown(` - বৈশ্বিক পরিসর (সবখানে পাবে)`);
                }
                md.appendMarkdown(`\n\n`);
                md.appendMarkdown(`**মান:** \`${displayValue}\``);
                return new vscode.Hover(md);
            }
            
            // Check if hovering over a map/module property (e.g., json.name or as.greet)
            const lineText = document.lineAt(position.line).text;
            const wordStart = position.character - word.length;
            const beforeWord = lineText.substring(0, wordStart);
            const mapPropertyMatch = beforeWord.match(/([a-zA-Z_][a-zA-Z0-9_]*)\.$/);

            if (mapPropertyMatch) {
                const mapName = mapPropertyMatch[1];

                // PRIORITY 1: Check if it's an import alias (ano "..." hisabe mapName)
                const aliasImportRegex = new RegExp(`ano\\s+["']([^"']+)["']\\s+hisabe\\s+${mapName}\\s*;?`, 'm');
                const aliasImport = aliasImportRegex.exec(text);

                if (aliasImport) {
                    const modulePath = aliasImport[1];
                    try {
                        const currentDir = path.dirname(document.uri.fsPath);
                        const fullPath = path.resolve(currentDir, modulePath);

                        if (!fs.existsSync(fullPath)) {
                            console.log(`[BanglaCode] Module file not found for hover: ${fullPath}`);
                            return null;
                        }

                        // Check if it's a JSON file
                        if (modulePath.endsWith('.json')) {
                            const jsonText = fs.readFileSync(fullPath, 'utf8');
                            try {
                                const jsonData = JSON.parse(jsonText);
                                if (jsonData.hasOwnProperty(word)) {
                                    const value = jsonData[word];
                                    const md = new vscode.MarkdownString();
                                    md.appendMarkdown(`**${mapName}.${word}**\n\n`);
                                    md.appendMarkdown(`📦 JSON থেকে: \`${modulePath}\`\n\n`);

                                    if (typeof value === 'object' && value !== null && !Array.isArray(value)) {
                                        md.appendMarkdown(`**টাইপ:** object\n\n`);
                                        md.appendMarkdown(`**প্রপার্টি:** ${Object.keys(value).join(', ')}\n\n`);
                                        md.appendMarkdown(`**ব্যবহার করো:** \`${mapName}.${word}.propertyName\``);
                                    } else if (Array.isArray(value)) {
                                        md.appendMarkdown(`**টাইপ:** array\n\n`);
                                        md.appendMarkdown(`**দৈর্ঘ্য:** ${value.length}`);
                                    } else {
                                        md.appendMarkdown(`**টাইপ:** ${typeof value}\n\n`);
                                        md.appendMarkdown(`**মান:** \`${JSON.stringify(value)}\``);
                                    }

                                    return new vscode.Hover(md);
                                }
                            } catch (err) {
                                console.error('[BanglaCode] Error parsing JSON for hover:', err);
                            }
                        } else if (modulePath.endsWith('.bang')) {
                            // It's a .bang module - check for ONLY exported functions (with 'pathao' keyword)
                            const moduleText = fs.readFileSync(fullPath, 'utf8');

                            // Get function signature (only exported functions with 'pathao')
                            const funcInModuleRegex = new RegExp(`pathao\\s+kaj\\s+${word}\\s*\\(([^)]*)\\)`, 'm');
                            const funcInModule = funcInModuleRegex.exec(moduleText);

                            if (funcInModule) {
                                const params = funcInModule[1].trim();
                                const paramList = params ? params.split(',').map(p => p.trim()) : [];
                                const docComment = extractDocComment(moduleText, word, 'kaj');
                                const returnValue = extractReturnValue(moduleText, word);

                                const md = new vscode.MarkdownString();
                                md.appendMarkdown(`**kaj ${word}(${params})**\n\n`);
                                md.appendMarkdown(`📦 ইম্পোর্ট করা মডিউল: \`${modulePath}\`\n\n`);
                                md.appendMarkdown(`🔗 অ্যালিয়াস: \`${mapName}\`\n\n`);

                                if (docComment) {
                                    md.appendMarkdown('---\n\n');
                                    md.appendMarkdown(docComment + '\n\n');
                                    md.appendMarkdown('---\n\n');
                                }

                                md.appendMarkdown(`**প্যারামিটার:** ${paramList.length === 0 ? 'নেই' : paramList.join(', ')}\n\n`);

                                if (returnValue) {
                                    md.appendMarkdown(`**রিটার্ন করে:** \`${returnValue}\`\n\n`);
                                } else {
                                    md.appendMarkdown(`**রিটার্ন করে:** কিছু না (void)\n\n`);
                                }

                                md.appendMarkdown(`**কল করো:** \`${mapName}.${word}(${paramList.map((_, i) => `arg${i+1}`).join(', ')})\``);

                                return new vscode.Hover(md);
                            } else {
                                console.log(`[BanglaCode] Function '${word}' not found in module ${modulePath}`);
                            }
                        }
                    } catch (err) {
                        console.error('[BanglaCode] Error reading module for hover:', err);
                    }
                } else {
                    // Check if it's a map/JSON variable property
                    const mapKeys = extractMapKeys(text, mapName);
                    const keyInfo = mapKeys.find(k => k.name === word);
                    if (keyInfo) {
                        const md = new vscode.MarkdownString();
                        md.appendMarkdown(`**${mapName}.${word}**\n\n`);
                        if (keyInfo.nested) {
                            md.appendMarkdown(`**টাইপ:** object\n\n`);
                            md.appendMarkdown(`**নেস্টেড প্রপার্টি:** ${keyInfo.nested.join(', ')}\n\n`);
                            md.appendMarkdown(`**নেস্টেড ব্যবহার করো:** \`${mapName}.${word}.propertyName\``);
                        } else {
                            md.appendMarkdown(`**টাইপ:** ${keyInfo.type}\n\n`);
                            md.appendMarkdown(`\`${mapName}\` অবজেক্টের প্রপার্টি`);
                        }
                        return new vscode.Hover(md);
                    }
                }
            }
            
            // Check if function is from a directly imported module (ano "..." without hisabe alias)
            const importRegex = /ano\s+["']([^"']+)["']\s*;/g;
            let importMatch;
            while ((importMatch = importRegex.exec(text)) !== null) {
                const modulePath = importMatch[1];

                // Skip if this import has an alias (we already handled those above)
                const hasAlias = new RegExp(`ano\\s+["']${modulePath.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')}["']\\s+hisabe`, 'm').test(text);
                if (hasAlias) {
                    continue;
                }

                // Try to get function signature from module
                try {
                    const currentDir = path.dirname(document.uri.fsPath);
                    const fullPath = path.resolve(currentDir, modulePath);

                    if (!fs.existsSync(fullPath)) {
                        console.log(`[BanglaCode] Module file not found for direct import hover: ${fullPath}`);
                        continue;
                    }

                    if (!modulePath.endsWith('.bang')) {
                        // Only .bang modules can export functions directly
                        continue;
                    }

                    const moduleText = fs.readFileSync(fullPath, 'utf8');

                    // Look for ONLY exported functions (with 'pathao' keyword)
                    const funcInModuleRegex = new RegExp(`pathao\\s+kaj\\s+${word}\\s*\\(([^)]*)\\)`, 'm');
                    const funcInModule = funcInModuleRegex.exec(moduleText);

                    if (funcInModule) {
                        const params = funcInModule[1].trim();
                        const paramList = params ? params.split(',').map(p => p.trim()) : [];
                        const docComment = extractDocComment(moduleText, word, 'kaj');
                        const returnValue = extractReturnValue(moduleText, word);

                        const md = new vscode.MarkdownString();
                        md.appendMarkdown(`**kaj ${word}(${params})**\n\n`);
                        md.appendMarkdown(`📦 ইম্পোর্ট করা মডিউল: \`${modulePath}\`\n\n`);

                        if (docComment) {
                            md.appendMarkdown('---\n\n');
                            md.appendMarkdown(docComment + '\n\n');
                            md.appendMarkdown('---\n\n');
                        }

                        md.appendMarkdown(`**প্যারামিটার:** ${paramList.length === 0 ? 'নেই' : paramList.join(', ')}\n\n`);

                        if (returnValue) {
                            md.appendMarkdown(`**রিটার্ন করে:** \`${returnValue}\`\n\n`);
                        } else {
                            md.appendMarkdown(`**রিটার্ন করে:** কিছু না (void)\n\n`);
                        }

                        md.appendMarkdown(`**কল করো:** \`${word}(${paramList.map((_, i) => `arg${i+1}`).join(', ')})\``);

                        return new vscode.Hover(md);
                    }
                } catch (err) {
                    console.error('[BanglaCode] Error reading module for direct import hover:', err);
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
        
        // Find all declared variables (dhoro, sthir, bishwo)
        const declaredVars = new Set();
        const varDeclRegex = /(dhoro|sthir|bishwo)\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*=/g;
        let match;
        while ((match = varDeclRegex.exec(text)) !== null) {
            declaredVars.add(match[2]);
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

            // Skip if it's a function definition or keyword
            if (funcName === 'kaj' || funcName === 'jodi' || funcName === 'jotokkhon' ||
                funcName === 'ghuriye' || funcName === 'shuru') {
                continue;
            }

            // Count arguments (simple approach - count commas + 1, unless empty)
            const argCount = argsStr ? argsStr.split(',').length : 0;

            // Check local functions first
            if (funcDefs.has(funcName)) {
                const def = funcDefs.get(funcName);

                if (argCount !== def.paramCount) {
                    const pos = document.positionAt(match.index);
                    const range = new vscode.Range(pos, pos.translate(0, funcName.length));
                    const paramNames = def.params ? def.params.split(',').map(p => p.trim()).join(', ') : '';
                    diagnostics.push(new vscode.Diagnostic(
                        range,
                        `Function '${funcName}(${paramNames})' expects ${def.paramCount} argument(s) but got ${argCount}`,
                        vscode.DiagnosticSeverity.Error
                    ));
                }
                continue; // Skip to next iteration
            }

            // Check built-in functions with known parameter requirements
            const builtinFunc = builtinFunctions.find(f => f.label === funcName);
            if (builtinFunc) {
                // Extract expected param count from insertText
                const paramMatch = builtinFunc.insertText.match(/\$\{(\d+)(?::([^}]+))?\}/g);
                if (paramMatch) {
                    const expectedCount = paramMatch.length;

                    if (argCount !== expectedCount) {
                        const pos = document.positionAt(match.index);
                        const range = new vscode.Range(pos, pos.translate(0, funcName.length));
                        diagnostics.push(new vscode.Diagnostic(
                            range,
                            `Built-in function '${funcName}' expects ${expectedCount} argument(s) but got ${argCount}`,
                            vscode.DiagnosticSeverity.Error
                        ));
                    }
                }
                continue; // Skip to next iteration
            }
        }
        
        // Check for aliased module function calls (e.g., helper.greet(name))
        // First, find all import aliases and their modules
        const importAliases = new Map();
        const aliasImportRegex = /ano\s+["']([^"']+)["']\s+hisabe\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*;?/g;
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
                    if (fs.existsSync(fullPath) && modulePath.endsWith('.bang')) {
                        const moduleText = fs.readFileSync(fullPath, 'utf8');

                        // Find function definition in module (ONLY exported with 'pathao')
                        const funcInModuleRegex = new RegExp(`pathao\\s+kaj\\s+${funcName}\\s*\\(([^)]*)\\)`);
                        const funcInModule = funcInModuleRegex.exec(moduleText);

                        if (funcInModule) {
                            const params = funcInModule[1].trim();
                            const paramNames = params.split(',').map(p => p.trim()).join(', ');
                            const expectedCount = params ? params.split(',').length : 0;
                            const actualCount = argsStr ? argsStr.split(',').length : 0;

                            if (actualCount !== expectedCount) {
                                const pos = document.positionAt(match.index);
                                const range = new vscode.Range(pos, pos.translate(0, aliasName.length + 1 + funcName.length));
                                diagnostics.push(new vscode.Diagnostic(
                                    range,
                                    `Function '${funcName}(${paramNames})' expects ${expectedCount} argument(s) but got ${actualCount}`,
                                    vscode.DiagnosticSeverity.Error
                                ));
                            }
                        } else {
                            // Function not found or not exported
                            const pos = document.positionAt(match.index);
                            const range = new vscode.Range(pos, pos.translate(0, aliasName.length + 1 + funcName.length));
                            diagnostics.push(new vscode.Diagnostic(
                                range,
                                `Function '${funcName}' is not exported from module '${modulePath}'`,
                                vscode.DiagnosticSeverity.Error
                            ));
                        }
                    }
                } catch (err) {
                    console.error('[BanglaCode] Diagnostic error for aliased import:', err);
                }
            }
        }

        // Check for directly imported functions (ano "..." without hisabe alias)
        const directImports = [];
        const directImportRegex = /ano\s+["']([^"']+)["']\s*;/g;
        while ((match = directImportRegex.exec(text)) !== null) {
            const modulePath = match[1];
            // Skip if this import has an alias
            const hasAlias = new RegExp(`ano\\s+["']${modulePath.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')}["']\\s+hisabe`, 'm').test(text);
            if (!hasAlias && modulePath.endsWith('.bang')) {
                directImports.push(modulePath);
            }
        }

        // For each directly imported module, check function calls
        for (const modulePath of directImports) {
            try {
                const currentDir = path.dirname(document.uri.fsPath);
                const fullPath = path.resolve(currentDir, modulePath);

                if (!fs.existsSync(fullPath)) continue;

                const moduleText = fs.readFileSync(fullPath, 'utf8');

                // Find all exported functions in this module
                const exportedFuncs = new Map();
                const exportedFuncRegex = /pathao\s+kaj\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)/g;
                let exportMatch;

                while ((exportMatch = exportedFuncRegex.exec(moduleText)) !== null) {
                    const funcName = exportMatch[1];
                    const params = exportMatch[2].trim();
                    const paramCount = params ? params.split(',').length : 0;
                    exportedFuncs.set(funcName, { params, paramCount });
                }

                // Now check if any function calls match these exported functions
                const funcCallRegex2 = /([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)/g;
                let callMatch;

                while ((callMatch = funcCallRegex2.exec(text)) !== null) {
                    const funcName = callMatch[1];
                    const argsStr = callMatch[2].trim();

                    // Skip keywords and built-ins
                    if (funcName === 'kaj' || funcName === 'jodi' || funcName === 'jotokkhon' ||
                        funcName === 'ghuriye' || funcName === 'shuru') {
                        continue;
                    }

                    // Skip if it's a local function
                    if (funcDefs.has(funcName)) continue;

                    // Skip if it's a built-in
                    if (builtinFunctions.some(f => f.label === funcName)) continue;

                    // Check if it's an exported function from this module
                    if (exportedFuncs.has(funcName)) {
                        const def = exportedFuncs.get(funcName);
                        const argCount = argsStr ? argsStr.split(',').length : 0;

                        if (argCount !== def.paramCount) {
                            const pos = document.positionAt(callMatch.index);
                            const range = new vscode.Range(pos, pos.translate(0, funcName.length));
                            const paramNames = def.params.split(',').map(p => p.trim()).join(', ');
                            diagnostics.push(new vscode.Diagnostic(
                                range,
                                `Function '${funcName}(${paramNames})' expects ${def.paramCount} argument(s) but got ${argCount}`,
                                vscode.DiagnosticSeverity.Error
                            ));
                        }
                    }
                }
            } catch (err) {
                console.error('[BanglaCode] Diagnostic error for direct import:', err);
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
