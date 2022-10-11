import axios from "axios";

const getVideos = async () => {
  axios.get(`/videos`).then((res) => res.data.result);
};

const getVideo = async (id: string | string[] | undefined) => {
  return axios
  .get(`/videos/${id}`)
  .then((res) => res.data.result)
  .catch((err) => null)
};

const getVideoPage = async (page:number, type:string | undefined) => {
  return axios
  .get(`/videos`,{
    params: {offset: page, categoryId: type??'1'}
  })
  .then((res) => res)
  .catch((err) => null)
};

const VideoAPI = {
  getVideos,
  getVideo,
  getVideoPage,
};

export default VideoAPI;
