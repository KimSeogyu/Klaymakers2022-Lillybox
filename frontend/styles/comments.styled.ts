import styled from "styled-components";

export const Container = styled.div`
  margin-left:20px;
`;
export const NewComments = styled.div`
  display: flex;
  align-items: center;
  gap: 10px;
`;
export const Avatar = styled.img`
  width: 50px;
  height: 50px;
  border-radius: 50%;
`;
interface Iinput {
  bottom: boolean,
};
export const Input = styled.input`
  border: none;
  background-color: transparent;
  outline: none;
  padding: 5px;
  // width: 100%;
  min-width: 450px;
  //margin: auto;
  ${(props:Iinput) =>
		props.bottom ? 
    'border-bottom: 1px solid;'
    : 'border-bottom: 1px solid soft;'
	}

`;
export const DivComments = styled.div`
  display: flex;
  gap: 30px;
  margin: 30px 0px;
`;
export const Details = styled.div`
  display: flex;
  flex-direction: column;
  gap: 10px;
  width: 60%;
  max-width: 60%;
`;
export const Name = styled.span`
  font-size: 13px;
  font-weight: 500;
`;
export const SpanDate = styled.span`
  font-size: 12px;
  font-weight: 400;
  margin-left: 5px;
`;
export const Text = styled.div`
  font-size: 14px;
`;
export const Button = styled.button`
  justify-self: flex-right;
  align-self: flex-start;
  text-align: right;
  padding: 10px 15px;
  border: none;
  border-radius: 3px;
  color: black;
  font-weight: 500;
  margin-top: 10px;
  margin-right: 20px;
  cursor: default;
  ${(props:any) =>
		props.color && `
    background-color: ${props.color};
    border: 1px solid ${props.color};
    cursor: pointer;
    `
	}
`;
