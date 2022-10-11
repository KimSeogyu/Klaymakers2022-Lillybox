import styled from "styled-components";

export const Modal = styled.div`
  display: flex;
  width: 100vw;
  height: 100vh;
  flex-flow: row-wrep;
  justify-content: center;
  align-items: center;
`;

export const ModalBackdrop = styled.div`
  width: 100%;
  height: 100%;
  position: fixed;
  display: flex;
  flex-flow: row wrep;
  justify-content: center;
  align-items: center;
  background: rgba(0, 0, 0, 0.5);
`;

export const ModalView = styled.div.attrs((props) => ({
  role: "dialog",
}))`
  width: 30%;
  padding: 1.5rem;
  background: white;
  border-radius: 2px;

  h3 {
    color: cornflowerblue;
    margin: 0;
    font-size: 1.5rem;
  }
  
  p {
    color: cornflowerblue;
    font-size: 1.125rem;
  }
`;

export const ButtonGroup = styled.div`
  margin-top: 1.5rem;
  display: flex;
  color: cornflowerblue;
  justify-content: flex-end;
`;

export const ModalButton = styled.button`
  display: flex;
  gap: 20px;
  margin-right: 20px;
  background: transparent;
  border-radius: 3px;
  border: 2px solid cornflowerblue;
  color: cornflowerblue;
  margin: 0 1em;
  padding: 0.25em 1em;
  &: hover {
    background-color: cornflowerblue;
    color: white;
  }
`;

export const InputField = styled.input`
  width: 100%;
  font-size: 18px;
  padding: 10px;
  padding-top: 10px;
  padding-bottom: 10px;
  margin-top: 10px;
  margin-bottom: 10px;
  background: white;
  border: 1px solid cornflowerblue;
  border-radius: 3px;
  color: black;
  ::placeholder {
    color: cornflowerblue;
  }
`;
