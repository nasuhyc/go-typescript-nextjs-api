// ** Reactstrap Imports
import {
  Col,
  Container,
  Row,
  Button,
  Card,
  CardBody,
  CardTitle,
  UncontrolledTooltip,
} from "reactstrap";

// ** Third Party Component Imports
import DataTable from "react-data-table-component";
import Swal from "sweetalert2";
import withReactContent from "sweetalert2-react-content";

// ** React Imports
import { FormEvent, useEffect, useState } from "react";

// ** Icon Imports
import { FaEdit, FaTrash } from "react-icons/fa";

// ** Config Imports
import axios from "@/configs/axiosConfig";

// ** Component Imports
import ModalComponent from "@/components/ModalComponent";

// ** Type Imports
import type { ColumnsType, UserType } from "@/types/user";

// ** Next Imports
import { GetServerSideProps } from "next";
import FilterHeader from "@/components/FilterHeader";

interface PropsType {
  users: UserType[];
}

const MySwal = withReactContent(Swal);

// ** Component  Users is used to list users
const Users = (props: PropsType) => {
  // ** Vars
  const columns: ColumnsType[] = [
    {
      name: "file",
      selector: "photo",
      sortable: true,
      cell: (row: UserType) => (
        <img
          src={`http://127.0.0.1:8000/storage/users/${row.file?.file}`}
          alt=""
          width={30}
          height={30}
        />
      ),
    },
    {
      name: "Name",
      selector: "first_name",
      sortable: true,
    },
    {
      name: "Last Name",
      selector: "last_name",
      sortable: true,
    },
    {
      name: "Email",
      selector: "email",
      sortable: true,
    },
    {
      name: "Age",
      selector: "age",
      sortable: true,
    },
    {
      name: "Action",
      selector: "action",
      sortable: true,
      cell: (row: UserType) => (
        <>
          <Button
            color="primary"
            onClick={() => {
              editHandler(row.id ?? 0);
              toggle();
            }}
          >
            <FaEdit />
          </Button>

          <Button
            color="danger"
            onClick={() => {
              deleteHandler(row.id ?? 0);
            }}
          >
            <FaTrash />
          </Button>
        </>
      ),
    },
  ];

  // ** Props
  const { users } = props;

  console.log({ users });

  // ** States
  const [modal, setModal] = useState(false);
  const [usersData, setUsers] = useState<UserType[]>([]);
  const [getUser, setGetUser] = useState<UserType | undefined>(undefined);

  const toggle = () => {
    if (!modal) {
      setGetUser(undefined);
    }

    setModal((modal) => !modal);
  };

  // ** UseEffect
  useEffect(() => {
    setUsers(users);
  }, []);

  // ** Functions

  //* Submit Handler for Create and Update
  const submitHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.target as HTMLFormElement);

    if (typeof getUser !== "undefined" && getUser.id) {
      updateHandler(getUser.id, formData);
    } else {
      createHandler(formData);
    }
  };

  //* Delete Handler for Delete User by Id
  const deleteHandler = async (id: number) => {
    MySwal.fire({
      title: "Are you sure?",
      text: "You cannot undo this action!",
      showCancelButton: true,
      confirmButtonColor: "#DD6B55",
      confirmButtonText: "Yes, Delete!",
    }).then(async (status) => {
      if (status.isConfirmed) {
        await axios.delete(`user/delete/${id}`).then(() => {
          setUsers(usersData.filter((user: any) => user.id !== id));
        });
      }
    });
  };

  //* Edit Handler for Get User by Id
  const editHandler = async (id: number) => {
    await axios.get(`user/get/${id}`).then((res) => {
      setGetUser(res.data.user);
    });
  };

  //* Create Handler for Create User
  const createHandler = async (formData: FormData) => {
    await axios.post("user/create", formData).then((res: any) => {
      modal && toggle();
      setUsers([...usersData, res.data.user]);
    });
  };

  //* Update Handler for Update User by Id
  const updateHandler = async (id: number, formData: FormData) => {
    await axios.put(`user/update/${id}`, formData).then((res) => {
      const usersShallow = [...usersData];

      const newUsers = usersShallow.map((user) => {
        if (user.id === id) {
          return res.data.user;
        }

        return user;
      });
      modal && toggle();
      setUsers(newUsers);
    });
  };

  //* Filter Handler for Filter Users
  const filterHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const formData = new FormData(e.target as HTMLFormElement);
    const query = new URLSearchParams(formData as URLSearchParams).toString();

    await axios.get(`user/getAll?${query}`).then((res) => {
      setUsers(res.data.data);
    });
  };

  return (
    <Container fluid className="p-5">
      <Row>
        <Card>
          <CardBody>
            <Row className="justify-content-between">
              <Col xs="6">
                <CardTitle tag="h5">User Page</CardTitle>
              </Col>
              <Col xs="6" className="text-end">
                <Button color="success" onClick={toggle}>
                  Add New User <i className="fa fa-plus"></i>
                </Button>
              </Col>
            </Row>

            <ModalComponent
              userData={getUser}
              modal={modal}
              toggle={toggle}
              submitHandler={submitHandler}
              getUser={getUser}
            />

            <h3 className="pt-5">User List</h3>
            <FilterHeader filterHandler={filterHandler} />

            <DataTable columns={columns as never} data={usersData} />
          </CardBody>
        </Card>
      </Row>
    </Container>
  );
};

export default Users;

// ** Get Server Side Props for Get All Users
export const getServerSideProps: GetServerSideProps = async () => {
  try {
    const request = await axios.get("user/getAll");

    const { data } = request.data;

    return {
      props: {
        users: data,
      },
    };
  } catch (error) {
    return {
      redirect: {
        permanent: true,
        destination: "/505",
      },
      props: {},
    };
  }
};
