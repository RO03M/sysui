import { useResourceInfo } from "../api/get-resource-info"
import { CpuProgressBar } from "./dashboard/cpu-progress-bar";

export function CpuCores() {
    const { cores } = useResourceInfo();

    return (
        <div>
            {cores.map((core) => (
                <CpuProgressBar
                    label={`${core.label}`}
                    percentage={core.usage}
                />
            ))}
        </div>
    )
}