export let replaced = input => {
  return input
    .replace(/\\\\/g, '!')
    .replace(/\\"/g, '@')
    .replace(/\\x../g, '$$')
};

export let numChars = (input:string):number => {
  return replaced(input).length - 2;
};

export let expand = (input:string): number => {
  return input
    .replace(/"/g, '!"')
    .replace(/\\/g, '!!')
    .length + 2;
};
