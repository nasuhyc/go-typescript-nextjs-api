// ** Reactstrap Imports
import { FormEvent } from "react";
import {
  FormGroup,
  Input,
  Label,
  Form,
  Button,
  Modal,
  ModalHeader,
  ModalBody,
  ModalFooter,
} from "reactstrap";

import { UserType } from "../types/user";

// ** Types & Interfaces for ModalComponent
type PropsType = {
  userData: UserType | undefined;
  modal: boolean;
  toggle: () => void;
  submitHandler: (e: FormEvent<HTMLFormElement>) => Promise<void>;
  getUser: UserType | undefined;
};

// ** Component ModalComponent is used to add and edit users
const ModalComponent = (props: PropsType) => {
  // ** Props
  const { userData, modal, toggle, submitHandler, getUser } = props;

  return (
    <Modal isOpen={modal} toggle={toggle}>
      <ModalHeader toggle={toggle}>Modal title</ModalHeader>
      <Form autoComplete="off" onSubmit={(e) => submitHandler(e)}>
        <ModalBody>
          <FormGroup>
            <Label for="photo">Photo</Label>
            <Input
              required
              type="file"
              name="file"
              id="file"
              placeholder="Lütfen Fotoğraf Yükleyiniz."
              // defaultValue={data.file.file}
            />
          </FormGroup>

          <FormGroup>
            <Label for="first_name">Name</Label>
            <Input
              type="text"
              name="first_name"
              id="firstName"
              placeholder="Lütfen İsminizi Giriniz."
              defaultValue={userData?.first_name}
            />
          </FormGroup>

          <FormGroup>
            <Label for="last_name">Surname</Label>
            <Input
              type="text"
              name="last_name"
              id="lastName"
              placeholder="Lütfen Soyisminizi Giriniz."
              defaultValue={userData?.last_name}
            />
          </FormGroup>

          <FormGroup>
            <Label for="email">Email</Label>
            <Input
              type="email"
              name="email"
              id="email"
              placeholder="Lütfen Email adresinizi Giriniz."
              defaultValue={userData?.email}
            />
          </FormGroup>

          <FormGroup>
            <Label for="age">Yaş</Label>
            <Input
              type="number"
              name="age"
              id="age"
              placeholder="Lütfen Email adresinizi Giriniz."
              defaultValue={userData?.age}
            />
          </FormGroup>
        </ModalBody>
        <ModalFooter>
          <Button color="secondary" onClick={toggle}>
            Back
          </Button>
          <Button type="submit" color="primary">
            {getUser ? "Güncelle" : "Kaydet"}
          </Button>
        </ModalFooter>
      </Form>
    </Modal>
  );
};

export default ModalComponent;
