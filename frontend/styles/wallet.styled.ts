import styled from "styled-components";

export const WalletWrapper = styled.div``;
export const WalletInfo = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
`;
export const Info = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 20px;
`;
export const Blocks = styled.div`
  display: grid;
  justify-content: center;
  align-items: center;
  grid-template-columns: repeat(8, 1fr);
  gap: 20px;
`;
export const Block = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  border-radius: 5px;
  min-width: 10%;
  gap: 10px;
  ${(props:any) =>
		props.color &&`
    background-color: ${props.color};
    `
	}
`;
export const BlockInfo = styled.span`
  display: flex;
  color: black;
`;
export const Hr = styled.hr`
  margin: 15px 0px;
  border: 0.5px solid;
  ${(props: any) => props.type && `display: none;`}
`;
export const Text = styled.div`
  flex-direction: column;
  text-align: center;
  gap: 10px;
  width: 60%;
  margin: 0 auto;
`;
export const Buttons = styled.div`
  display: flex;
  gap: 20px;
  margin-right: 20px;
  :hover{
    transition:all 0.9s;
    cursor:pointer;
    color:orange;
  }
`;
export const Button = styled.button`
  justify-self: flex-right;
  align-self: flex-start;
  text-align: center;
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