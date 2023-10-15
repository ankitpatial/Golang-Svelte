/*
 * Copied from: https://github.com/emerald-code/slug-generator/blob/master/index.js
 * Author: Ankit Patial
 * Dated:  09/05/22, 3:19 PM
 */


const map: any = {
	// latin
	À: 'A',
	Á: 'A',
	Â: 'A',
	Ã: 'A',
	Ä: 'A',
	Å: 'A',
	Æ: 'AE',
	Ç: 'C',
	È: 'E',
	É: 'E',
	Ê: 'E',
	Ë: 'E',
	Ì: 'I',
	Í: 'I',
	Î: 'I',
	Ï: 'I',
	Ð: 'D',
	Ñ: 'N',
	Ò: 'O',
	Ó: 'O',
	Ô: 'O',
	Õ: 'O',
	Ö: 'O',
	Ő: 'O',
	Ø: 'O',
	Ù: 'U',
	Ú: 'U',
	Û: 'U',
	Ü: 'U',
	Ű: 'U',
	Ý: 'Y',
	Þ: 'TH',
	ß: 'ss',
	à: 'a',
	á: 'a',
	â: 'a',
	ã: 'a',
	ä: 'a',
	å: 'a',
	æ: 'ae',
	ç: 'c',
	è: 'e',
	é: 'e',
	ê: 'e',
	ë: 'e',
	ì: 'i',
	í: 'i',
	î: 'i',
	ï: 'i',
	ð: 'd',
	ñ: 'n',
	ò: 'o',
	ó: 'o',
	ô: 'o',
	õ: 'o',
	ö: 'o',
	ő: 'o',
	ø: 'o',
	ù: 'u',
	ú: 'u',
	û: 'u',
	ü: 'u',
	ű: 'u',
	ý: 'y',
	þ: 'th',
	ÿ: 'y',
	ẞ: 'SS',
	// serbian latin
	š: 's',
	đ: 'dj',
	ž: 'z',
	č: 'c',
	ć: 'c'
};

export default function slugify(str: string) {
	if (!str) return '';

  str = str.toUpperCase();
	const replacement = '_';

	return (
    str
			.split('')
			.reduce(function (result: string, char: string) {
				if (map[char]) {
					char = map[char];
				}
				char = char.replace(/[^\w\s$\-]/g, '');
				result += char;
				return result;
			}, '')
			// trim leading/trailing spaces
			.replace(/^\s+|\s+$/g, '')
			// convert spaces
			.replace(/[-\s]+/g, replacement || '-')
			// remove trailing separator
			.replace('#{replacement}$', '')
	);
}
