/* eslint-disable react-hooks/exhaustive-deps */
import VideoAPI from "./videos";
import { useInfiniteQuery } from "react-query";
import type { IInfinite } from "./types";
import { useEffect, useState } from "react";
import { useTheme } from "next-themes";
import {
	BsFillSunFill,
	BsFillMoonFill,
} from "react-icons/bs";

export const ThemeSwitch = () => {
	const [mounted, setMounted] = useState(false);
	const { theme, setTheme } = useTheme();
	useEffect(() => {
	  if (!theme)
		setTheme("dark");
	  setMounted(true);
	}, []);
  
	if (!mounted) {
	  return null;
	}
	return (
		<div
		  className="relative mr-3 md:mr-7"
		  onClick={() => setTheme(theme === "dark" ? "light" : "dark")}
		>
		  {theme === "light" ? <BsFillSunFill /> : <BsFillMoonFill />}
		</div>
	);
};

export const useInfiniteScrollQuery = (category: string) => {
	const getPageVideo = async ({ pageParam = 0 }) => {
	  const res = await VideoAPI.getVideoPage(pageParam, category);
	  return {
		// 실제 데이터
		result: res ? res.data.result : null,
		// 반환 값에 현재 페이지를 넘겨주자
		current_page: pageParam,
		// 페이지가 마지막인지 알려주는 서버에서 넘겨준 true/false 값
		isLast:res &&  res.data.result.length < 20 ? true : false,
	  };
	};
	const {
	  data: getVideos,
	  fetchNextPage: getNextPage,
	  isSuccess: getVideosIsSuccess,
	  hasNextPage: getNextPageIsPossible,
	} = useInfiniteQuery<IInfinite, Error>(["page_video_list"], getPageVideo, {
	  getNextPageParam: (lastPage, pages) => {
		if (!lastPage.isLast) return lastPage.current_page + 1;
		return null;
	  },
	});
	return { getVideos, getNextPage, getVideosIsSuccess, getNextPageIsPossible };
};

export const isInteger = (obj: string | number): boolean => {
    console.log(Number.isInteger(Number(obj)));
    return Number.isInteger(Number(obj));
  };

