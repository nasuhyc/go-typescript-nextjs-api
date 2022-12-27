// ** Reactstrap Imports
import { FormEvent } from "react";
import { Col, Input, Row, Button } from "reactstrap";

// ** Types & Interfaces for FilterHeader
interface PropsType {
  filterHandler: (e: FormEvent<HTMLFormElement>) => Promise<void>;
}

// ** Component FilterHeader is used to filter users
const FilterHeader = (props: PropsType) => {
  // ** Props
  const { filterHandler } = props;

  return (
    <Col className="col-12">
      <form onSubmit={(e) => filterHandler(e)}>
        <Row className="col-12 pt-3 mb-3 align-items-center">
          <Col className="col-4">
            <Input
              type="text"
              placeholder="First Name"
              name="first_name"
              className="col-5"
            />
          </Col>
          <Col className="col-4">
            <Input
              type="text"
              placeholder="Last Name"
              name="last_name"
              className="col-5"
            />
          </Col>
          <Col className="col-4 text-end">
            <Button type="submit">Filter</Button>
          </Col>
        </Row>
      </form>
    </Col>
  );
};

export default FilterHeader;
