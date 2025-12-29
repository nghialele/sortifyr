import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { convertPlaylistDuplicates, convertPlaylists, convertPlaylistUnplayables } from "../types/playlist"
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

export const usePlaylistGetUnplayables = () => {
  return useQuery({
    queryKey: ["playlist", "unplayable"],
    queryFn: async () => (await apiGet(`${ENDPOINT}/unplayable`, convertPlaylistUnplayables)).data,
    staleTime: STALE_TIME.MIN_30,
    throwOnError: true,
  })
}

export const usePlaylistRemoveDuplicates = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: () => apiPost(`${ENDPOINT}/duplicate`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["task"] })
      queryClient.invalidateQueries({ queryKey: ["task_history"] })
    },
  })
}
