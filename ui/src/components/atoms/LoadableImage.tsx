import { ComponentProps, useState } from "react";
import { LoadingSpinner } from "../molecules/LoadingSpinner";
import { cn } from "@/lib/utils";

type Props = ComponentProps<"img">

export const LoadableImage = ({ className, ...props }: Props) => {
  const [loaded, setLoaded] = useState(false)

  return (
    <div className="relative w-full h-full">
      {!loaded && (
        <div className="absolute inset-0 flex items-center justify-center">
          <LoadingSpinner />
        </div>
      )}

      <img
        onLoad={() => setLoaded(true)}
        className={cn("w-full h-full object-cover transition-opacity duration-300", loaded ? "opacity-100" : "opacity-0", className)}
        {...props}
      />
    </div>
  )
}
