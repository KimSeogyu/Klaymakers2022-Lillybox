import type { AppProps } from "next/app";
import { useEffect, useState } from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import Head from "next/head";
import Sidebar from "../components/Sidebar";
import Header from "../components/Header";
import Image from "next/image";
import "../styles/globals.css";
import { ThemeProvider } from "next-themes";
const queryClient = new QueryClient();

function MyApp({ Component, pageProps }: AppProps) {
  const [isOpen, setIsOpen] = useState(true);
  const toggleSidebar = () => setIsOpen(!isOpen);
  const [isMounted, setIsMounted] = useState(false);
  useEffect(() => {
    setTimeout(() => {
      setIsMounted(true);
    }, 2000);
  }, []);
  return (
    <>
      <Head>
        <link rel="icon" href="/lilly-plain.png" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <title>Lillybox</title>
      </Head>
      <QueryClientProvider client={queryClient}>
      <ThemeProvider attribute="class">
      <div className={isMounted ? "block" : "hidden"}>
        <Header toggleSidebar={toggleSidebar} isMounted={isMounted} />
        <div className="grid grid-cols-7 mt-20">
          <Sidebar isOpen={isOpen} />
          <div className="col-span-6 relative min-h-screen z-0">
            <div className="ml-4 lg:ml-17">
              <Component {...pageProps} />
            </div>
          </div>
        </div>
      </div>
      </ThemeProvider>
      <div
          className={
            isMounted
              ? "hidden"
              : "grid place-items-center min-h-screen text-red-500"
          }
        >
          <Image src="/lilly-yellow-box.png" width={450} height={450} alt="" />
        </div>
      </QueryClientProvider>
    </>
  );
}

export default MyApp;
