import { useQuery } from "@tanstack/react-query"
import { convertPlaylists } from "../types/playlist"
import { STALE_TIME } from "../types/staletime"
import { apiGet } from "./query"

const ENDPOINT = "playlist"

export const usePlaylistGetAll = () => {
  return useQuery({
    queryKey: ["playlist"],
    queryFn: async () => (await apiGet(ENDPOINT, convertPlaylists)).data,
    staleTime: STALE_TIME.MIN_30,
    throwOnError: true,
  })
}
