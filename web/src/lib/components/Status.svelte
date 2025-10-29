<script lang="ts">
    import { onDestroy, onMount } from 'svelte';

    let { status = 'out' } = $props();

    let interval: number = 0;

    let time = $state('the time :)');
    let date = $state('the date :)');

    let out = $derived(status == 'out');
    let free = $derived(status == 'free');
    let busy = $derived(status == 'busy');
    let working = $derived(status == 'working');

    let statusText = $derived.by(() => {
        switch (status) {
            case 'free':
                return 'Chilling';
            case 'busy':
                return 'In a meeting';
            case 'working':
                return "Head down don't bother me";
            case 'out':
            default:
                return 'Out of Office';
        }
    });

    let i = 0;
    const updateTimes = () => {
        const now = new Date();
        time = now.toLocaleTimeString();
        date = now.toLocaleDateString();

        status = ['out', 'free', 'busy', 'working'][(i = (i + 1) % 4)];
    };

    onMount(() => {
        updateTimes;
        interval = setInterval(updateTimes, 500);
    });

    onDestroy(() => {
        clearInterval(interval);
    });
</script>

<div class="component-container">
    <div class="status-container">
        <div class={['circle', { out, free, working, busy }]}></div>
        <p>{statusText}</p>
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
            background: black;
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
    }
</style>
