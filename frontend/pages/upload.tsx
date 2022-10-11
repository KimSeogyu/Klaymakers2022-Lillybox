import React, { useRef, useState } from "react";
import { BiCloud, BiPlus } from "react-icons/bi";
import Link from "next/link";
import Caver from "caver-js";
import ipfsClient from "ipfs-http-client";
import { useRouter } from "next/router";
import abi from "../contract/abi/Lillybox.json";
import Image from "next/image";

export default function Upload() {
	// Creating state for the input field
	const [title, setTitle] = useState("");
	const [description, setDescription] = useState("");
	const [category, setCategory] = useState("");
	const [location, setLocation] = useState("");
	const [thumbnail, setThumbnail] = useState<File>();
	const [video, setVideo] = useState<File>();

	//  Creating a ref for thumbnail and video
	const thumbnailRef = useRef() as React.MutableRefObject<HTMLInputElement>;
	const videoRef = useRef() as React.MutableRefObject<HTMLInputElement>;

	const router = useRouter();

	const handleSubmit = async ({
		title,
		description,
		category,
		thumbnail,
		video,
	}: any) => {
		if (!window.klaytn) {
			throw new Error(
				"Please install kaikas!\nURL: https://chrome.google.com/webstore/detail/kaikas/jblndlipeogpafnldhgmapagcccfchpi"
			);
		}

		if (!video || !thumbnail) {
			throw new Error("비디오와 썸네일을 모두 올려주세요");
		}

		const caver = new Caver(window.klaytn);
		const [account] = await window.klaytn.enable();
		const projectId = `${process.env.NEXT_PUBLIC_PROJECT_ID}`;
		const projectSecret = `${process.env.NEXT_PUBLIC_PROJECT_SECRET}`;
		const client = ipfsClient({
			host: "ipfs.infura.io",
			port: 5001,
			protocol: "https",
			headers: {
				authorization:
					"Basic " +
					Buffer.from(projectId + ":" + projectSecret).toString("base64"),
			},
		});

		const thumbnailCid = await client.add(thumbnail);
		const videoCid = await client.add(video);
		console.log(thumbnailCid);
		console.log(videoCid);
		const metadata = {
			metadata: {
				title: "Token Metadata",
				type: "object",
				properties: {
					name: title,
					type: "videos",
					description: description,
					categories: [category],
					video_uri: `https://lillybox.infura-ipfs.io/ipfs/${videoCid.path}`,
					thumbnail_uri: `https://lillybox.infura-ipfs.io/ipfs/${thumbnailCid.path}`,
					created_at: new Date(),
				},
			},
		};

		const metadataCid = await client.add(JSON.stringify(metadata));
		const contract = new caver.klay.Contract(
			abi as any,
			`${process.env.NEXT_PUBLIC_CONTRACT_ADDR}`
		);
		const transactionReceipt = await contract.send(
			{
				from: account,
				gas: 500000,
			},
			"mintVod",
			metadataCid.path
		);
		alert("Success");
		await router.push("/");
	};

	return (
		<div className="flex flex-row">
			<div className="flex-1 flex flex-col">
				<div className="mt-5 mr-10 flex justify-end">
					<div className="flex items-center">
						<Link href="/">
							<button className="bg-transparent  text-[#9CA3AF] py-2 px-6 border rounded-lg  border-gray-600  mr-6">
								Discard
							</button>
						</Link>
						<button
							onClick={() => {
								handleSubmit({
									title,
									description,
									category,
									thumbnail,
									video,
								});
							}}
							className="bg-blue-500 hover:bg-blue-700 text-white  py-2  rounded-lg flex px-4 justify-between flex-row items-center"
						>
							<BiCloud />
							<p className="ml-2">Upload</p>
						</button>
					</div>
				</div>
				<div className="flex flex-col m-10     mt-5  lg:flex-row">
					<div className="flex lg:w-3/4 flex-col ">
						<label>Title</label>
						<input
							value={title}
							onChange={(e) => setTitle(e.target.value)}
							placeholder="Rick Astley - Never Gonna Give You Up (Official Music Video)"
							className="w-[90%] placeholder:text-gray-600  rounded-md mt-2 h-12 p-2 border border-[#444752] focus:outline-none"
						/>
						<label>Description</label>
						<textarea
							value={description}
							onChange={(e) => setDescription(e.target.value)}
							placeholder="Never Gonna Give You Up was a global smash on its release in July 1987, topping the charts in 25 countries including Rick’s native UK and the US Billboard Hot 100.  It also won the Brit Award for Best single in 1988. Stock Aitken and Waterman wrote and produced the track which was the lead-off single and lead track from Rick’s debut LP “Whenever You Need Somebody."
							className="w-[90%] h-32 placeholder:text-gray-600  rounded-md mt-2 p-2 border  border-[#444752] focus:outline-none"
						/>

						<div className="flex flex-row mt-10 w-[90%] justify-between">
							<div className="flex flex-col w-2/5">
								<label>Location</label>
								<input
									value={location}
									onChange={(e) => setLocation(e.target.value)}
									type="text"
									placeholder="Bali - Indonesia"
									className="w-[90%] placeholder:text-gray-600  rounded-md mt-2 h-12 p-2 border border-[#444752] focus:outline-none hover:bg-yellow-200 yellow:hover:bg-yellow-200"
								/>
							</div>
							<div className="flex flex-col w-2/5">
								<label>Category</label>
								<select
									value={category}
									onChange={(e) => setCategory(e.target.value)}
									className="w-[90%] placeholder:text-gray-600  rounded-md mt-2 h-12 p-2 border border-[#444752] focus:outline-none hover:bg-yellow-200 yellow:hover:bg-yellow-200"
								>
									<option>Music</option>
									<option>Sports</option>
									<option>Gaming</option>
									<option>News</option>
									<option>Entertainment</option>
									<option>Education</option>
									<option>Science & Technology</option>
									<option>Travel</option>
									<option>Other</option>
								</select>
							</div>
						</div>
						<label className="text-[#9CA3AF]  mt-10 text-sm">Thumbnail</label>

						<div
							onClick={() => {
								thumbnailRef.current.click();
							}}
							className="border-2 w-64 border-gray-600  border-dashed rounded-md mt-2 p-2  h-36 items-center justify-center flex hover:bg-yellow-200 yellow:hover:bg-yellow-200"
						>
							{thumbnail ? (
								<div style={{width: '100%', height: '100%', position: 'relative'}}>
                  <Image
                    layout="fill"
                    onClick={() => {
                      thumbnailRef.current.click();
                    }}
                    src={URL.createObjectURL(thumbnail)}
                    alt="thumbnail"
                    className="h-full rounded-md"
                  />
								</div>
							) : (
								<BiPlus size={40} color="gray" />
							)}
						</div>

						<input
							type="file"
							className="hidden"
							ref={thumbnailRef}
							onChange={(e) => {
								setThumbnail((e.target.files as FileList)[0]);
							}}
						/>
					</div>

					<div
						onClick={() => {
							videoRef.current.click();
						}}
						className={
							video
								? " w-96   rounded-md  h-64 items-center justify-center flex"
								: "border-2 border-gray-600  w-96 border-dashed rounded-md mt-8   h-64 items-center justify-center flex hover:bg-yellow-200 yellow:hover:bg-yellow-200"
						}
					>
						{video ? (
							<video
								controls
								src={URL.createObjectURL(video)}
								className="h-full rounded-md"
							/>
						) : (
							<p className="text-[#9CA3AF] hover:backdrop-blur-0">
								Upload Video
							</p>
						)}
					</div>
				</div>
				<input
					type="file"
					className="hidden"
					ref={videoRef}
					accept={"video/*"}
					onChange={(e) => {
						setVideo((e.target.files as FileList)[0]);
						console.log((e.target.files as FileList)[0]);
					}}
				/>
			</div>
		</div>
	);
}
