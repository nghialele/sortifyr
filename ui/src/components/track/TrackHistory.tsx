import { useTrackGetHistory } from "@/lib/api/track";
import { TrackHistoryFilter } from "@/lib/types/track";
import { formatDate } from "@/lib/utils";
import { DatesRangeValue } from '@mantine/dates';
import { useState } from "react";
import { SectionTitle } from "../atoms/Page";
import { DatePickerInput } from "../molecules/DatePickerInput";
import { Table } from "../molecules/Table";

export const TrackHistory = () => {
  const [range, setRange] = useState<[Date | null, Date | null]>([null, null])
  const [filter, setFilter] = useState<TrackHistoryFilter>({})

  const handleRangeChange = (r: DatesRangeValue) => {
    setRange(r)

    if (r[0] && r[1] || (!r[0] && !r[1])) {
      r[0]?.setHours(0)
      r[1]?.setHours(23)
      r[1]?.setMinutes(59)
      r[1]?.setSeconds(59)

      setFilter({ ...filter, start: r[0] ?? undefined, end: r[1] ?? undefined })
    }
  }

  const { history, isLoading, isFetchingNextPage, hasNextPage, fetchNextPage } = useTrackGetHistory(filter)

  const handleBottom = () => {
    if (!hasNextPage) return
    if (isFetchingNextPage) return

    fetchNextPage()
  }

  return (
    <>
      <SectionTitle
        title="Recently Played"
        description={`An overview of recently played tracks.\nRun the tracks task if a track has no title.`}
      />
      <DatePickerInput
        type="range"
        allowSingleDateInRange
        placeholder="Filter by date range"
        value={range}
        onChange={handleRangeChange}
      />
      <Table
        idAccessor="historyId"
        columns={[
          { accessor: "name", title: "Track", width: "80%", ellipsis: true },
          { accessor: "playedAt", render: ({ playedAt }) => <p className="text-muted">{formatDate(playedAt)}</p> }
        ]}
        records={history}
        noRecordsText="No recorded tracks yet"
        fetching={isLoading || isFetchingNextPage}
        onScrollToBottom={handleBottom}
      />
    </>
  )
}
