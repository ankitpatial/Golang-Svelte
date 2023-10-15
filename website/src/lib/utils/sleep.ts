/**
 * Function to sleep a thread for sometime
 */
export default function sleep(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}
