import React, { useState } from "react";
import ReactPlayer from "react-player";
import Comments from "../../components/Comments";
import { useQuery, useMutation } from "react-query";
import { useRouter } from "next/router";
import {
  Container,
  Content,
  VideoWrapper,
  Title,
  Details,
  Info,
  Buttons,
  Hr,
  Description,
} from "../../styles/video.styled";
import axios from "axios";
import {
  ButtonGroup,
  ModalBackdrop,
  ModalView,
  InputField,
  ModalButton,
} from "../../styles/modal.styled";
import { callDontaion, setApprovalForAll } from "../../lib/contract";
import type { ILillyVideo } from "../../lib/types";
import VideoAPI from "../../lib/videos";


const WatchVideoPage = () => {
  const router = useRouter();
  const { id } = router.query;
  const { isLoading, isError, error, data } = useQuery<ILillyVideo, Error>(
    `/videos/${id}`,
    () => VideoAPI.getVideo(id)
  );
  const [donateModalOpen, setDonateModalOpen] = useState(false);
  const [donateAmount, setDonateAmount] = useState(0);
  const [LILAmount, setLILAmount] = useState(0);
  const donateModalHandler = () => {
    setDonateModalOpen(!!!donateModalOpen);
  };
  const onChangeDonateAmount = (e: any) => {
    setDonateAmount(e.target.value);
  };
  const onChangeLILAmount = (e: any) => {
    setLILAmount(e.target.value);
  };
  const donate = async () => {
    if (window.klaytn === undefined) {
      alert(
        "Please install kaikas!\nURL: https://chrome.google.com/webstore/detail/kaikas/jblndlipeogpafnldhgmapagcccfchpi"
      );
    } else {
      const approval = await setApprovalForAll();
      const account = data?.account;
      if (account == null) {
        alert(`NULL Account`);
      }
      const receipt = await callDontaion(account, donateAmount, LILAmount);
      alert(JSON.stringify(receipt, null, 2));
    }
  };

  const addDeclareation = useMutation((newDeclareation) =>
    axios.post(`/videos/${id}`, newDeclareation)
  );
  if (isLoading)
    return (
      <div>
        <h1>Loading...</h1>
      </div>
    );
  if (error)
    return (
      <div>
        <h1>Loading...{error.message}</h1>
      </div>
    );
  if (!data)
    return (
      <div>
        <h1>No data..</h1>
      </div>
    );

  return (
    <>
      {!isError ? (
        <Container>
          <Content>
            <VideoWrapper>
              <ReactPlayer
                url={data?.video_uri}
                playing={true}
                controls={true}
                loop={true}
                muted={true}
                playsinline={true}
                style={{ width: "100%", height: "100%" }}
              />
            </VideoWrapper>
            <Details>
              <Info>
                {data.categories &&
                  data.categories?.map((type: string, idx) => {
                    return <p key={idx}>#{type} </p>;
                  })}
              </Info>
            </Details>
            <Title>{data?.name}</Title>
            <Details>
              <Info>
                <p>{data?.views} views</p>
                <p> • </p>
                <p>
                  {new Date(data.created_at).getFullYear()}.
                  {new Date(data?.created_at).getMonth()}.
                  {new Date(data?.created_at).getDate()}{" "}
                </p>
                <Buttons onClick={donateModalHandler}>Donation</Buttons>
              </Info>
            </Details>
            <Hr />
            <Title>{data?.nickname}</Title>
            <Description>{data?.description}</Description>
            <Hr />
            <Comments props={data} />
          </Content>
          <Hr style={{ display: "none" }} />
          {donateModalOpen ? (
            <ModalBackdrop>
              <ModalView>
                <h3>Donation</h3>
                <p>Donation Amount</p>
                <InputField
                  type="text"
                  placeholder="amount"
                  onChange={onChangeDonateAmount}
                />
                <p>LIL Amount</p>
                <InputField
                  type="text"
                  placeholder="amount"
                  onChange={onChangeLILAmount}
                />
                <ButtonGroup>
                  <ModalButton onClick={donateModalHandler}>Cancle</ModalButton>
                  <ModalButton onClick={donate}>Confirm</ModalButton>
                </ButtonGroup>
              </ModalView>
            </ModalBackdrop>
          ) : null}
        </Container>
      ) : (
        <div>
          <h1>an error has occurred!</h1>
          <p>{error}</p>
        </div>
      )}
    </>
  );
};

export default WatchVideoPage;
