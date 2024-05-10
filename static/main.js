const kanaWordElement = document.getElementById('kana-word');
const inputField = document.getElementById('input-field');
const statusElement = document.getElementById('status');
const scoreElement = document.getElementById('score');

let currentWord = '';
let score = 0;

function generateWord() {
    const hiraganaChars = 'あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんがぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぷぺぽぁぃぅぇぉゃゅょっ';
    const katakanaChars = 'アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヰヱヲンガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ';

    const kanaChars = kanaType === 'hiragana' ? hiraganaChars : (kanaType === 'katakana' ? katakanaChars : hiraganaChars + katakanaChars);
    const wordLength = Math.floor(Math.random() * 5) + 1;
    let word = '';

    for (let i = 0; i < wordLength; i++) {
        word += kanaChars[Math.floor(Math.random() * kanaChars.length)];
    }

    currentWord = word;
    kanaWordElement.textContent = currentWord;
    inputField.value = '';
    statusElement.textContent = '';
}

function handleInput(event) {
    if (event.key === 'Enter') {
        const input = inputField.value.toLowerCase();
        const romaji = toRomaji(currentWord);

        if (input === romaji) {
            statusElement.textContent = '🎉 Correct!';
            score++;
            scoreElement.textContent = score;
            generateWord();
        } else {
            statusElement.textContent = `😭 Incorrect. The correct Romaji for '${currentWord}' is '${romaji}'.`;
        }
    }
}

function toRomaji(word) {
    const kanaMap = {
        'あ': "a", 'い': "i", 'う': "u", 'え': "e", 'お': "o",
        'か': "ka", 'き': "ki", 'く': "ku", 'け': "ke", 'こ': "ko",
        'さ': "sa", 'し': "shi", 'す': "su", 'せ': "se", 'そ': "so",
        'た': "ta", 'ち': "chi", 'つ': "tsu", 'て': "te", 'と': "to",
        'な': "na", 'に': "ni", 'ぬ': "nu", 'ね': "ne", 'の': "no",
        'は': "ha", 'ひ': "hi", 'ふ': "fu", 'へ': "he", 'ほ': "ho",
        'ま': "ma", 'み': "mi", 'む': "mu", 'め': "me", 'も': "mo",
        'や': "ya", 'ゆ': "yu", 'よ': "yo",
        'ら': "ra", 'り': "ri", 'る': "ru", 'れ': "re", 'ろ': "ro",
        'わ': "wa", 'を': "o", 'ん': "n",
        'が': "ga", 'ぎ': "gi", 'ぐ': "gu", 'げ': "ge", 'ご': "go",
        'ざ': "za", 'じ': "ji", 'ず': "zu", 'ぜ': "ze", 'ぞ': "zo",
        'だ': "da", 'ぢ': "ji", 'づ': "zu", 'で': "de", 'ど': "do",
        'ば': "ba", 'び': "bi", 'ぶ': "bu", 'べ': "be", 'ぼ': "bo",
        'ぱ': "pa", 'ぴ': "pi", 'ぷ': "pu", 'ぺ': "pe", 'ぽ': "po",
        'ぁ': "a", 'ぃ': "i", 'ぅ': "u", 'ぇ': "e", 'ぉ': "o",
        'ゃ': "ya", 'ゅ': "yu", 'ょ': "yo", 'っ': "tsu",
        'ア': "a", 'イ': "i", 'ウ': "u", 'エ': "e", 'オ': "o",
        'カ': "ka", 'キ': "ki", 'ク': "ku", 'ケ': "ke", 'コ': "ko",
        'サ': "sa", 'シ': "shi", 'ス': "su", 'セ': "se", 'ソ': "so",
        'タ': "ta", 'チ': "chi", 'ツ': "tsu", 'テ': "te", 'ト': "to",
        'ナ': "na", 'ニ': "ni", 'ヌ': "nu", 'ネ': "ne", 'ノ': "no",
        'ハ': "ha", 'ヒ': "hi", 'フ': "fu", 'ヘ': "he", 'ホ': "ho",
        'マ': "ma", 'ミ': "mi", 'ム': "mu", 'メ': "me", 'モ': "mo",
        'ヤ': "ya", 'ユ': "yu", 'ヨ': "yo",
        'ラ': "ra", 'リ': "ri", 'ル': "ru", 'レ': "re", 'ロ': "ro",
        'ワ': "wa", 'ヰ': "i", 'ヱ': "e", 'ヲ': "o", 'ン': "n",
        'ガ': "ga", 'ギ': "gi", 'グ': "gu", 'ゲ': "ge", 'ゴ': "go",
        'ザ': "za", 'ジ': "ji", 'ズ': "zu", 'ゼ': "ze", 'ゾ': "zo",
        'ダ': "da", 'ヂ': "ji", 'ヅ': "zu", 'デ': "de", 'ド': "do",
        'バ': "ba", 'ビ': "bi", 'ブ': "bu", 'ベ': "be", 'ボ': "bo",
        'パ': "pa", 'ピ': "pi", 'プ': "pu", 'ペ': "pe", 'ポ': "po",
    };

    let romaji = '';
    for (const char of word) {
        romaji += kanaMap[char] || '';
    }
    return romaji;
}

generateWord();
inputField.addEventListener('keydown', handleInput);
