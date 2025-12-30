import { useMutation, useQueryClient } from "@tanstack/react-query"
import { apiPost, NO_CONVERTER, NO_DATA } from "./query"

const ENDPOINT = "setting"

export const useSettingUploadExport = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (file: File) => apiPost(`${ENDPOINT}/export`, NO_DATA, NO_CONVERTER, [{ field: "zip", file }]),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["task"] })
      queryClient.invalidateQueries({ queryKey: ["task_history"] })
    }
  })
}
