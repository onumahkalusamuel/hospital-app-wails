export function useDebounce () {
  let timeout = null as never as NodeJS.Timeout;
  return function(fnc: Function, delayMs?: number) {
    clearTimeout(timeout);
    timeout = setTimeout(() => { fnc(); }, delayMs || 300);
  };
}