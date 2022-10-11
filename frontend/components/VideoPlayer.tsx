import React, { useEffect, useState } from "react";
import Image from "next/image";
import Link from "next/link";
import { useInView } from "react-intersection-observer";
import type { ILillyVideo } from "../lib/types";
import timeForToday from "../lib/timeForToday";
import { useStore } from "../lib/store";
import shallow from 'zustand/shallow';
import { Genres } from "./data/Genres";
import { useInfiniteScrollQuery } from "../lib/utils";

function VideoPlayer() {
  const { ref, inView } = useInView();
  const { category } = useStore((state) => ({ category: state.category }), shallow)
  const [id, setId] = useState(Genres[0].id);  
  const { getVideos, getNextPage, getVideosIsSuccess, getNextPageIsPossible } =
    useInfiniteScrollQuery(id);
  useEffect(() => {
    setId(category);
    if (inView && getNextPageIsPossible && getVideosIsSuccess) {
      getNextPage();
    }
  }, [
    inView,
    getVideos,
    getNextPage,
    category,
    getNextPageIsPossible,
    getVideosIsSuccess,
  ]);

  return (
    <div>
      {getVideosIsSuccess && getVideos?.pages ? (
        <>
          <section className="video-section">
            {" "}
            {getVideos.pages.map((page) => {
              const result = page.result;
              return (
                result &&
                result.map((item: ILillyVideo) => {
                  return (
                    <article
                      className="video-container"
                      ref={ref}
                      key={item.id}
                    >
                      <Link
                        href={{
                          pathname: `/watch/[id]`,
                          query: { id: item.id },
                        }}
                        className="thumbnail"
                        passHref
                        prefetch={false}
                      >
                        <a>
                          <Image
                            unoptimized
                            className="thumbnail-image"
                            src={item.thumbnail_uri}
                            alt="thumbnail image"
                            width="250"
                            height="150"
                          ></Image>
                        </a>
                      </Link>
                      <div className="video-details">
                        <Link
                          href={{
                            pathname: "/watch/[id]",
                            query: { id: item.id },
                          }}
                        >
                          <a
                            className="video-title text-ellipsis overflow-hidden"
                            aria-label={item.name}
                            title={item.name}
                          >
                            {item.name}
                          </a>
                        </Link>
                        <Link
                          href={{
                            pathname: "/watch/[id]",
                            query: { id: item.id },
                          }}
                        >
                          <a
                            className="video-channel-name text-ellipsis overflow-hidden"
                            aria-label={item.nickname}
                            title={item.nickname}
                          >
                            {item.nickname}
                          </a>
                        </Link>
                        <div className="video-metadata">
                          <span>{`${item.views} views`}</span>
                          <span> • </span>
                          <span>{timeForToday(item.created_at)}</span>
                        </div>
                      </div>
                    </article>
                  );
                })
              );
            })}
          </section>
          <div className="flex justify-center ...">
            <div> • </div>
          </div>
        </>
      ) : (
        <>
          <section className="video-section" />
          <div className="flex justify-center">
            <div> No videos </div>
          </div>
        </>
      )}
    </div>
  );
}

export default VideoPlayer;
