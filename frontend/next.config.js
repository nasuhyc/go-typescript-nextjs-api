/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  env: {
    BACKEND_API: "http://127.0.0.1:8000/api/",
  },
  images: {
    domains: ["http://127.0.0.1:8000"],
  },
};

module.exports = nextConfig;
