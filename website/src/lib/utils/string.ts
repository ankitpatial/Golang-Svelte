// humanize a string
export function humanize(str: string) {
  if (!str) return '';

  // return str
  //     .replace(/^[\s_]+|[\s_]+$/g, '')
  //     .replace(/^[a-z]/, function (m) {
  //         return m.toUpperCase();
  //     });

  return str
    .replace(/[_\s]+/g, ' ')
    .replace(/([A-Z]+)/g, ' $1')
    .replace(/([A-Z][a-z])/g, ' $1')
}


