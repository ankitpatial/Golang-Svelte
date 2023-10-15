/*
 * Author: Ankit Patial
 */

export function copyToClipboard(input: HTMLInputElement) {
  return () => {
    if (!navigator?.clipboard) return;

    input.select();
    input.setSelectionRange(0, 99999); // For mobile devices
    navigator.clipboard.writeText(input.value);
  };
}
