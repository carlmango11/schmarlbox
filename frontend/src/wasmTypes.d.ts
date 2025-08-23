declare global {
  export interface Window {
    Go: any;
    getDisplay: () => string[];
  }
}

export {};
