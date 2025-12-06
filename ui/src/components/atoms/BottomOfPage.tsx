import { cn } from "@/lib/utils";
import { LoadingSpinner } from "../molecules/LoadingSpinner";

interface Props {
  showLoading?: boolean;
  hasNextPage?: boolean;
  ref?: React.Ref<HTMLDivElement>;
}

export function BottomOfPage({ showLoading = false, hasNextPage = true, ref }: Props) {
  return (
    <div className={cn(hasNextPage && "sticky left-0 h-24")} ref={ref}>
      {showLoading && <LoadingSpinner />}
    </div>
  );
}
