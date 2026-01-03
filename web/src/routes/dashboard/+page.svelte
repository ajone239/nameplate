<script lang="ts">
    import JobTitle from '$lib/components/JobTitle.svelte';
    import Status from '$lib/components/Status.svelte';
    import { onMount } from 'svelte';

    let interval: number = 0;

    let message = $state('');
    let status = $state('Away');

    const updateMessage = async () => {
        const res = await fetch('/api/data');
        const { message: message1 } = await res.json();

        message = message1;
    };

    const updateStatus = async () => {
        const res = await fetch('/api/status');
        const { status: new_status } = await res.json();

        status = new_status;
    };

    const updateAll = async () => {
        await updateMessage();
        await updateStatus();
    };

    onMount(async () => {
        await updateAll();
        interval = setInterval(updateAll, 500);
    });
</script>

<div class="dash-container">
    <div class="whoami-container">
        <h1 class="name-container">Austin Jones</h1>

        <JobTitle />
    </div>

    <Status {status} />

    <p>Hold for theme toggle component</p>

    <p>Hold for high score component</p>

    <p>Hold for Wordle component</p>

    <div class="bottom"><p>test api hit: [{message}]</p></div>
</div>

<style>
    * {
        border: red 1px solid;
    }

    .dash-container {
        margin: 5px;
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        grid-template-rows: 1fr, 1fr, 20px;
    }

    .whoami-container {
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        grid-column: 1 / 3;
        grid-row: 1 / 2;
    }

    .name-container {
        font-size: 4.5em;
        border-bottom: white 3px solid;
        margin: 10px;
    }

    .bottom {
        display: flex;
        justify-content: center;
        grid-column: 1 / 4;
        grid-row: 3 / 4;
    }
</style>
