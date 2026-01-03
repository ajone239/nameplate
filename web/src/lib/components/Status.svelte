<script lang="ts">
    import { onDestroy, onMount } from 'svelte';

    let { status = 'out' } = $props();

    let interval: number = 0;

    let time = $state('the time :)');
    let date = $state('the date :)');

    let out = $derived(status == 'Away');
    let free = $derived(status == 'Free');
    let busy = $derived(status == 'Busy');
    let meeting = $derived(status == 'Meeting');
    let working = $derived(status == 'Headdown');

    const updateTimes = async () => {
        const now = new Date();
        time = now.toLocaleTimeString();
        date = now.toLocaleDateString();
    };

    onMount(async () => {
        await updateTimes();
        interval = setInterval(updateTimes, 500);
    });

    onDestroy(() => {
        clearInterval(interval);
    });
</script>

<div class="component-container">
    <div class="status-container">
        <div class={['circle', { out, free, meeting, working, busy }]}></div>
        <p>{status}</p>
    </div>

    <div class="time-container">
        <p>{time}</p>
    </div>

    <div class="time-container">
        <p>{date}</p>
    </div>
</div>

<style>
    .component-container {
        background: grey;
        border-radius: 10px;
        margin: 5px;
        padding: 10px;
    }
    .status-container {
        display: flex;
        align-items: center;
        justify-content: center;
        border-bottom: white 1px solid;
    }
    .time-container {
        display: flex;
        justify-content: space-around;
        flex-direction: row;
    }
    .circle {
        margin-right: 10px;
        width: 20px;
        height: 20px;
        border-radius: 10px;
        background: grey;
        box-shadow: 0 0px 1px 1px lightgrey;

        &.out {
            background: grey;
            box-shadow: 0 0px 1px 1px lightgrey;
        }
        &.working {
            background: orange;
            box-shadow: 0 0px 1px 1px orangered;
        }
        &.busy {
            background: red;
            box-shadow: 0 0px 1px 1px pink;
        }
        &.free {
            background: green;
            box-shadow: 0 0px 1px 1px lightgreen;
        }
        &.meeting {
            background: red;
            box-shadow: 0 0px 1px 1px pink;
        }
    }
</style>
