export declare type View = {
  id: number;
  resourceId: number;
  resourceType: string;
  displayType: string;
  name: string;
  desc: string;
};

export declare type Resource = {
  id?: number;
  name: string;
  type: string;
  data: string;
};

export declare type Option = {
  id?: number;
  columns: Column[];
  parameters: Paramter[];
  version: number;
};

export declare type Paramter = {
  label: string;
  name: string;
  // type?: string; // todo only input now
};
export declare type Column = {
  label: string;
  name: string;
};

export declare type ResourceType = {
  value: string;
  label: string;
  color: string;
};

export const ResourceTypeValue: Array<ResourceType> = [
  { value: 'mysql', label: 'MySQL', color: 'blue' },
];

// export declare type DataType = {
//   value: string;
//   label: string;
//   color: string;
// };

// export const DataTypeValues: Array<DataType> = [
//   { value: 'string', label: 'String', color: 'blue' },
//   { value: 'number', label: 'Number', color: 'red' },
// ];
//
// export declare type Dict = {
//   id: number;
//   name: string;
//   details: Array<DictDetail>;
// };

// export declare type DictDetail = {
//   id: number;
//   key: string;
//   value: string;
// };
