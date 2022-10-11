import Link from 'next/link';
import { Items } from "./data/SidebarItems";

function Sidebar({ isOpen }: { isOpen: boolean }) {
  return (
    <div
      className={
        isOpen
          ? "flex flex-col justify-between  ml-0.5 mr-0.5 col-span-0.5 z-10 shadow-sm md:ml-10 md:mr-0"
          : "hidden"
      }
    >
      <ul className="sticky top-14 flex flex-col justify-between gap-10 overflow-y-scroll h-[90%]">
        {Items &&
          Items.map((item, index) => {
            return (
            <Link href={item.link ? `${item.link}` : '/'} key={index}>
              <li className="flex items-center text-center gap-4 transition-none p-3 cursor-pointer hover:text-gray-600 md:p-2">
                {item.icon}{" "}
                <span className="font-semibold pr-4 hidden lg:block">
                  {item.name}
                </span>
              </li>
            </Link>
            );
          })}
      </ul>
    </div>
  );
}

export default Sidebar;
