/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: false,
  images: {
    domains: ["picsum.photos"],
  },
  crossOrigin: "anonymous",
  async rewrites() {
    return [
      {
        source: "/request_id",
        destination: `${process.env.AUTH_SERVER_ENDPOINT}/request_id`,
      },
      {
        source: "/login",
        destination: `${process.env.AUTH_SERVER_ENDPOINT}/login`,
      },
      {
        source: "/sign",
        destination: `${process.env.AUTH_SERVER_ENDPOINT}/sign`,
      },
      {
        source: "/user",
        destination: `${process.env.AUTH_SERVER_ENDPOINT}/user`,
      },
      {
        source: "/videos",
        destination: `${process.env.NEXT_PUBLIC_API_URL}/videos`,
      },
      {
        source: "/videos/:path*",
        destination: `${process.env.NEXT_PUBLIC_API_URL}/videos/:path*`,
      },
    ];
  },
  webpack: (config, { isServer }) => {
    if (!isServer) {
      config.resolve.fallback.fs = false;
    }
    return config;
  },
};

module.exports = nextConfig;
