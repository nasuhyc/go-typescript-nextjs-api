import { FormGroup, Label, Input } from "reactstrap";

// ** Types & Interfaces
type PropsType = {
  id?: any;
  name: string;
  text: string;
  type: any;
  value?: any;
};

// ** Component DynamicFormGroup is used to create a form group dynamically
const DynamicFormGroup = (props: PropsType) => {
  const { name, text, type, value, id } = props;

  return (
    <FormGroup>
      <Label for={name}>{text}</Label>
      <Input
        type={type}
        name={name}
        id={id}
        placeholder={text}
        defaultValue={value}
      />
    </FormGroup>
  );
};
export default DynamicFormGroup;
