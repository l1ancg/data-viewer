export declare type Location = {
  offset: number;
  line: number;
  column: number;
};

export declare type ParseResult = {
  sql?: string;
  columnList?: string[];
  whereList?: string[];
  error?: { start: Location; end: Location };
};
