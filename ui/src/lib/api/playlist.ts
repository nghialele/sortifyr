import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { convertPlaylists } from "../types/playlist"
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

export const usePlaylistSync = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: () => apiPost(`${ENDPOINT}/sync`),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["playlist"] }),
    throwOnError: true,
  })
}
