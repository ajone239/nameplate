<script lang="ts">
    import { onDestroy, onMount } from 'svelte';

    let interval: number;

    const JobTitles = [
        'Software Engineer',
        'Resident Young Person',
        'Hipper than Alec',
        '3F Champion',
        'The Web Guy',
        "Formerly Keith's Cubemate",
        'The guardian of the Grassy Knoll',
        "Jeff's most loyal and trusted advisor"
    ];

    let jobIndex = $state(0);
    let text = $derived(JobTitles[jobIndex]);

    const typewriter = (node: HTMLElement, { speed = 1 }) => {
        const valid =
            node.childNodes.length === 1 && node.childNodes[0].nodeType === Node.TEXT_NODE;

        if (!valid) {
            throw new Error(`This transition only works on elements with a single text node child`);
        }

        const text = node.textContent;
        const duration = text.length / (speed * 0.01);

        return {
            duration,
            tick: (t: number) => {
                const i = Math.trunc(text.length * t);
                node.textContent = text.slice(0, i);
            }
        };
    };

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
    {#key jobIndex}
        <p in:typewriter={{ speed: 2 }} class="title">
            {text}
        </p>
    {/key}
</div>

<style>
    .title-container {
        display: flex;
        justify-content: flex-end;
        height: 20px;
        margin: 10px;
    }
    .title {
        all: unset;
        font-size: 1.85em;
        font-style: italic;
        color: grey;
    }
</style>
