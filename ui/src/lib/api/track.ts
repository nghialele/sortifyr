import { useInfiniteQuery } from "@tanstack/react-query"
import { STALE_TIME } from "../types/staletime"
import { convertTrackHistories } from "../types/track"
import { apiGet } from "./query"

const ENDPOINT = "track"
const PAGE_LIMIT = 100

export const useTrackGetHistory = () => {
  const { data, isLoading, fetchNextPage, isFetchingNextPage, hasNextPage, error, refetch, isFetching } = useInfiniteQuery({
    queryKey: ["track", "history"],
    queryFn: async ({ pageParam = 0 }) => {
      const queryParams = new URLSearchParams({
        page: pageParam.toString(),
        limit: PAGE_LIMIT.toString(),
      })

      const url = `${ENDPOINT}/history?${queryParams.toString()}`
      return (await apiGet(url, convertTrackHistories)).data
    },
    initialPageParam: 0,
    getNextPageParam: (lastPage, allPages) => {
      return lastPage.length < PAGE_LIMIT ? undefined : allPages.length + 1
    },
    staleTime: STALE_TIME.MIN_5,
    throwOnError: true,
  })

  const history = data?.pages.flat() ?? []

  return {
    history,
    isLoading,
    fetchNextPage,
    isFetchingNextPage,
    hasNextPage,
    error,
    refetch,
    isFetching,
  }
}
