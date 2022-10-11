import { Html, Head, Main, NextScript } from "next/document";
import React from 'react';

export default function Document() {
	return (
		<Html lang='en'>
			<Head>
				<meta name="author" content="lilly0" />
				<meta name="description" content="Lillybox, enjoy everything on web3" />
				<meta name="keywords" content="Lillybox, enjoy everything on web3" />
			</Head>
			<body className="bg-white text-black dark:bg-black dark:text-white">
				<div id='root'>
					<Main />
					<NextScript />
				</div>
			</body>
		</Html>
	)
} 