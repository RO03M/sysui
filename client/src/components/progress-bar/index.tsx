import styles from "./progress-bar.module.css";

interface ProgressBarProps {
    percentage: number;
}

export function ProgressBar(props: ProgressBarProps) {
    const { percentage = 0 } = props;

    return (
        <div className={styles.outer}>
            <div
                className={styles.bar}
                style={{ width: `${percentage}%` }}
            />
        </div>
    );
}