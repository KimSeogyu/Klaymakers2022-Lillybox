import {
  BsGear,
  BsHouse,
  BsWallet2,
  BsGraphUp,
  BsCollectionPlay,
  BsController,
  BsFilm,
  BsClockHistory,
  BsHandThumbsUp,
  BsLightbulb,
  BsMusicPlayer,
  BsTrophy
} from "react-icons/bs";

export const Items = [
  {
    name: "Home",
    icon: <BsHouse size={25} />,
  },
  {
    name: "Wallet",
    icon: <BsWallet2 size={25} />,
    link: "/wallet",
  },
  {
    name: "Trending",
    icon: <BsGraphUp size={25} />,
  },
  {
    name: "Subscriptions",
    icon: <BsCollectionPlay size={25} />,
  },
  {
    name: "Gaming",
    icon: <BsController size={25} />,
  },
  {
    name: "Films",
    icon: <BsFilm size={25} />,
  },
  {
    name: "History",
    icon: <BsClockHistory size={25} />,
  },
  {
    name: "Likes",
    icon: <BsHandThumbsUp size={25} />,
  },
  {
    name: "Learning",
    icon: <BsLightbulb size={25} />,
  },
  {
    name: "Sports",
    icon: <BsTrophy size={25} />,
  },
  {
    name: "Music",
    icon: <BsMusicPlayer size={25} />,
  },
  {
    name: "Settings",
    icon: <BsGear size={25} />,
  },
];
