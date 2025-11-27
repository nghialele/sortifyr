import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { apiGet, apiPost } from "./query"
import { convertDirectories, DirectorySchema } from "../types/directory"
import { STALE_TIME } from "../types/staletime"

const ENDPOINT = "directory"

export const useDirectoryGetAll = () => {
  return useQuery({
    queryKey: ["directory"],
    queryFn: async () => (await apiGet(ENDPOINT, convertDirectories)).data,
    staleTime: STALE_TIME.MIN_30,
    throwOnError: true,
  })
}

export const useDirectorySync = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (directories: DirectorySchema[]) => apiPost(`${ENDPOINT}/sync`, directories),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["directory"] }),
    throwOnError: true,
  })
}
