import React, { useState } from "react";
import { useQuery } from "react-query";
import type { ILillyVideo, IComment } from "../lib/types";
import timeForToday from "../lib/timeForToday";
import {
  DivComments,
  Container,
  NewComments,
  Details,
  SpanDate,
  Input,
  Name,
  Button,
} from "../styles/comments.styled";
import axios from "axios";
import { userInfoStore } from "../lib/store";
import { useRouter } from "next/router";

const Comments = ({ props }: { props: ILillyVideo }) => {
  const router = useRouter();
  const id = props.id;
  const { myAccount, myNickname } = userInfoStore();
  const [comment, setComment] = useState<IComment[]>();
  const [feedComment, setfeedComment] = useState<string>("");
  const [isValid, setIsValid] = useState(false);
  
  const { isLoading, isError, error, data } = useQuery(
    `/videos/${id}/comments`,
    async () => {
      await axios
      .get(`/videos/${id}/comments`)
      .then((res) => setComment(res.data.result))
      .catch((error) => console.log(error));
    },
    {
      cacheTime: 3000,
      refetchOnMount: false,
      refetchOnWindowFocus: false,
      onSuccess: (data) => {
        console.log("[comments] get데이터 요청 성공", data);
      },
    }
  );
  const createComment = () => {
    const createData: IComment = {
      description: feedComment,
      nickname: myNickname,
    };
    console.log(createData)
    axios
      .post(`/videos/${id}/comments`, createData)
      .then()
      .catch((error) => console.log(error));
  };

  const failToGetAccount = () => {
    alert("Please use it after Login");
    router.push('/');
  };
  console.log(myAccount, myNickname);
  if (isLoading)
    return (
      <div>
        <h1>Loading...</h1>
      </div>
    );
  if (error)
    return (
      <div>
        <h1>error...</h1>
      </div>
    );

  return (
    <Container>
      {!isError ? (
        <>
          <NewComments>
            <form>
              {myNickname} : 
              <Input
                type="text"
                value={feedComment}
                onChange={(e) => {
                  setfeedComment(e.target.value);
                }}
                onKeyUp={(e) => {
                  (e.target as HTMLTextAreaElement).value.length > 0
                    ? setIsValid(true)
                    : setIsValid(false);
                }}
                bottom={isValid}
                placeholder="Add a comment..."
              />
              <Button
                onClick={(e) => {
                  myAccount === "" ? failToGetAccount() : createComment();
                }}
                color={feedComment ?? "".length > 0 ? "orange" : "gray"}
                disabled={isValid ? false : true}
              >
                COMMENT
              </Button>
            </form>
          </NewComments>
          {comment ? (
            comment.map((item) => {
              return (
                <DivComments key={item.id}>
                  <Details>
                    <Name>
                      {item.nickname}
                      <SpanDate>
                        {item.created_at && timeForToday(item.created_at)}
                      </SpanDate>
                      <div>{item.description}</div>
                    </Name>
                  </Details>
                </DivComments>
              );
            })
          ) : (
            <div>
              <br />
              No comments.{" "}
            </div>
          )}
        </>
      ) : (
        <div>
          <h2>an error has occurred!</h2>
        </div>
      )}
    </Container>
  );
};

export default Comments;
