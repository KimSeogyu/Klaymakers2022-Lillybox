import React from "react";
import { useRouter } from "next/router";
import { Genres } from "./data/Genres";
import { useStore } from "../lib/store";

function Navbar() {
  const router = useRouter();
  const { setCategory } = useStore();
  return (
    <div className="sticky top-14 bg-white dark:bg-black flex overflow-x-scroll items-center gap-6 scrollbar-hide items-center justify-center border-b-2 pb-4 z-20">
      {Genres &&
        Genres.map((item, index) => {
          return (
            <button
              key={index}
              className="border-[#0000001a] border-2 rounded-[16px] bg-[#0000000d] pl-2 pr-2 pt-1 pb-1 cursor-pointer last:mr-24 hover:bg-gray-200 active:bg-gray-900 active:text-white"
              onClick={() => {
                setCategory(item.id);
              }}
            >
              {item.name}
            </button>
          );
        })}
    </div>
  );
}

export default Navbar;
