//* User type for the user table
export type UserType = {
  id?: number;
  first_name: string;
  last_name: string;
  email: string;
  age: string;
  file: FileType;
};
// * File type for the user table
type FileType = {
  ID?: number;
  file: string;
};

// * Column type for the table component
export type ColumnsType = {
  name: string;
  selector?: string;
  sortable: boolean;
  cell?: any;
};
