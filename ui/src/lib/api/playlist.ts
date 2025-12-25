import { useMutation, useQuery } from "@tanstack/react-query"
import { convertPlaylistDuplicates, convertPlaylists } from "../types/playlist"
import { STALE_TIME } from "../types/staletime"
import { apiGet, apiPost } from "./query"

const ENDPOINT = "playlist"

export const usePlaylistGetAll = () => {
  return useQuery({
    queryKey: ["playlist"],
    queryFn: async () => (await apiGet(ENDPOINT, convertPlaylists)).data,
    staleTime: STALE_TIME.MIN_30,
    throwOnError: true,
  })
}

export const usePlaylistGetDuplicates = () => {
  return useQuery({
    queryKey: ["playlist", "duplicate"],
    queryFn: async () => (await apiGet(`${ENDPOINT}/duplicate`, convertPlaylistDuplicates)).data,
    staleTime: STALE_TIME.MIN_30,
    throwOnError: true,
  })
}

export const usePlaylistRemoveDuplicates = () => {
  // No need to invalidate queries as the task will probably takes a while and is done in the background
  return useMutation({
    mutationFn: () => apiPost(`${ENDPOINT}/duplicate`),
    throwOnError: true,
  })
}
