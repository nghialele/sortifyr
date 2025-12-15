import { Page, PageTitle, Section, SectionTitle } from "@/components/atoms/Page";
import { Select } from "@/components/molecules/Select";
import { TrackAddedTable } from "@/components/track/TrackAddedTable";
import { TrackDeletedTable } from "@/components/track/TrackDeletedTable";
import { TrackHistoryTable } from "@/components/track/TrackHistoryTable";
import { usePlaylistGetAll } from "@/lib/api/playlist";
import { TrackFilter } from "@/lib/types/track";
import { scrollTo } from "@/lib/utils";
import { Button, Group } from "@mantine/core";
import { useState } from "react";

export const Tracks = () => {
  const { data: playlists, isLoading: isLoadingPlaylists } = usePlaylistGetAll()

  const [filterCreated, setFilterCreated] = useState<TrackFilter>({})
  const [filterDeleted, setFilterDeleted] = useState<TrackFilter>({})

  return (
    <Page>
      <Group justify="space-between">
        <PageTitle
          title="Tracks"
          description="All track related tools."
        />
        <Group>
          <Button onClick={() => scrollTo("track-played")} radius="lg">
            Played
          </Button>
          <Button onClick={() => scrollTo("track-added")} radius="lg">
            Added
          </Button>
          <Button onClick={() => scrollTo("track-deleted")} radius="lg">
            Removed
          </Button>
        </Group>
      </Group>

      <Section id="track-played" className="min-h-full">
        <SectionTitle
          title="Recently Played"
          description="An overview of recently played tracks."
        />
        <TrackHistoryTable />
      </Section>

      <Section id="track-added" className="min-h-full">
        <SectionTitle
          title="Recently Added"
          description="An overview of recently added tracks to playlists."
        />
        <Select
          data={playlists?.map(p => ({ value: p.id.toString(), label: p.name }))}
          value={filterCreated.playlistId?.toString()}
          onChange={(v) => setFilterCreated({ ...filterCreated, playlistId: v ? v : undefined })}
          placeholder="Filter track by playlist..."
          disabled={isLoadingPlaylists}
        />
        <TrackAddedTable filter={filterCreated} />
      </Section>

      <Section id="track-deleted" className="min-h-full">
        <SectionTitle
          title="Recently Deleted"
          description="An overview of recently deleted tracks from playlists."
        />
        <Select
          data={playlists?.map(p => ({ value: p.id.toString(), label: p.name }))}
          value={filterDeleted.playlistId?.toString()}
          onChange={(v) => setFilterDeleted({ ...filterDeleted, playlistId: v ? v : undefined })}
          placeholder="Filter track by playlist..."
          disabled={isLoadingPlaylists}
        />
        <TrackDeletedTable filter={filterDeleted} />
      </Section>
    </Page>
  )
}
