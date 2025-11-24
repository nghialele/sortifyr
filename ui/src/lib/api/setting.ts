import { useQuery } from "@tanstack/react-query"
import { apiGet } from "./query"
import { convertSetting } from "../types/setting"
import { STALE_TIME } from "../types/staletime"

const ENDPOINT = "setting"

export const useSettingGet = () => {
  return useQuery({
    queryKey: ["setting"],
    queryFn: async () => (await apiGet(ENDPOINT, convertSetting)).data,
    staleTime: STALE_TIME.MIN_30,
    throwOnError: true,
  })
}
