import { ConfigContext } from "expo/config";

export default ({ config }: ConfigContext) => {
  config.extra = {
    ...config.extra,
    NEWS_APP_API_KEY: process.env.NEWS_APP_API_KEY,
  };
  return config;
};
