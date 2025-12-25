import { usePlaylistGetDuplicates, usePlaylistRemoveDuplicates } from "@/lib/api/playlist"
import { Playlist } from "@/lib/types/playlist"
import { Track } from "@/lib/types/track"
import { Button, Group } from "@mantine/core"
import { useDisclosure } from "@mantine/hooks"
import { sortBy } from "lodash"
import { type DataTableSortStatus } from "mantine-datatable"
import { useMemo, useState } from "react"
import { SectionTitle } from "../atoms/Page"
import { Confirm } from "../molecules/Confirm"
import { Table } from "../molecules/Table"
import { PlaylistCover } from "./PlaylistCover"
import { notifications } from "@mantine/notifications"

interface Duplicate extends Track {
  amount: number;
}

export const PlaylistDuplicates = () => {
  const { data: playlists, isLoading } = usePlaylistGetDuplicates()
  const removeDuplicates = usePlaylistRemoveDuplicates()

  const [sortStatus, setSortStatus] = useState<DataTableSortStatus<Playlist>>({
    columnAccessor: "name",
    direction: "asc",
  })
  const records = useMemo(() => {
    const data = sortBy(playlists, sortStatus.columnAccessor);
    return sortStatus.direction === "desc" ? data.reverse() : data;
  }, [playlists, sortStatus])

  const [opened, { open, close }] = useDisclosure()

  const handleRemove = () => {
    removeDuplicates.mutate(undefined, {
      onSuccess: () => notifications.show({ variant: "success", message: "Duplicates are getting removed. Come back later to see the result", }),
      onSettled: () => close(),
    })
  }

  const duplicates: Record<number, Duplicate[]> = useMemo(() => {
    return Object.fromEntries(
      playlists?.map(p => {
        const counts = p.duplicates.reduce<Record<number, Duplicate>>((acc, t) => {
          acc[t.id] ??= { ...t, amount: 0 }
          acc[t.id].amount++
          return acc
        }, {})

        return [
          p.id,
          Object.values(counts).sort((a, b) => b.amount - a.amount),
        ]
      }) ?? []
    )
  }, [playlists])

  return (
    <>
      <Group justify="space-between">
        <SectionTitle
          title="Playlist duplicates"
          description={`Playlists with duplicate tracks\nClick on a row to see the duplicate tracks`}
        />
        <Button onClick={open} radius="lg" color="secondary.1" c="black">
          Remove duplicates
        </Button>
      </Group>
      <Table
        idAccessor="id"
        columns={[
          {
            accessor: "id",
            title: "",
            width: 52,
            render: playlist => <PlaylistCover playlist={playlist} />,
          },
          { accessor: "name", sortable: true },
          { accessor: "trackAmount", sortable: true },
          { accessor: "owner.name", sortable: true },
          {
            accessor: "duplicates",
            sortable: true,
            render: ({ duplicates }) => <p>{duplicates.length}</p>
          },
        ]}
        rowExpansion={{
          content: ({ record: { id } }) => (
            <Table
              noHeader
              backgroundColor="background.1"
              columns={[
                { accessor: "name" },
                { accessor: "amount" }
              ]}
              records={duplicates[id]}
              height={180}
              className="m-4"
            />
          )
        }}
        records={records}
        sortStatus={sortStatus}
        onSortStatusChange={setSortStatus}
        fetching={isLoading}
        animated={false}
      />

      <Confirm
        opened={opened}
        onClose={close}
        modalTitle="Duplicates"
        title="Remove Duplicates"
        description="Are you sure you want to remove all duplicates?"
        onConfirm={handleRemove}
      />
    </>
  )
}
