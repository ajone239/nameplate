<script lang="ts">
    import { onDestroy, onMount } from 'svelte';

    let interval: number;

    const JobTitles = [
        'Software Engineer',
        'Resident Young Person',
        'Hipper than Alec',
        '3F Champion',
        'The Web Guy',
        "Keith's Cubemate",
        'The guardian of the Grassy Knoll',
        "Jeff's most loyal and trusted advisor"
    ];

    let jobIndex = $state(0);
    let text = $derived(JobTitles[jobIndex]);

    const nextJob = () => {
        jobIndex = (jobIndex + 1) % JobTitles.length;
    };

    onMount(() => {
        interval = setInterval(nextJob, 5000);
    });

    onDestroy(() => {
        clearInterval(interval);
    });
</script>

<div class="title-container">
    <h4 class="title">
        {text}
    </h4>
</div>

<style>
    .title-container {
        display: flex;
        justify-content: flex-end;
        margin: 10px;
    }
    .title {
        all: unset;
        font-size: 1.85em;
        font-style: italic;
        color: grey;
    }
</style>
