<script>
	import Text from "../resuable/Text.svelte";
    import IconPatient from '~icons/mdi/patient';
    import IconInpatient from '~icons/medical-icon/i-inpatient';
    import IconInpatient24px from '~icons/healthicons/inpatient-24px';
    import IconPatientBandAlt from '~icons/healthicons/patient-band-alt';
    import IconPatientBand from '~icons/healthicons/patient-band';
    import IconOutpatient from '~icons/healthicons/outpatient';
    import IconOutpatientMaterial from '~icons/material-symbols/outpatient';

    const ratings = [
        {icon: IconPatientBand, text: "Active", percentage: "100%", bg: "#047064"},
        {icon: IconOutpatient, text: "Books", percentage: "100%", bg: "#160094"},
        {icon: IconPatient, text: "Patients", percentage: "100%", bg: "#960000"},
        {icon: IconPatient, text: "Patients", percentage: "100%", bg: "#960000"},
        {icon: IconPatient, text: "Patients", percentage: "100%", bg: "#960000"},
        {icon: IconPatient, text: "Patients", percentage: "100%", bg: "#960000"}
    ]
</script>


<section class="flex w-72 gap-5 mt-10">
    {#each ratings as item}
        <div class="text-center">
            <div class={`w-20 h-20 shadow-md rounded-full flex items-center justify-center ${item.bg}`} >
                <item.icon class="text-5xl"/>           
            </div>
            <Text text = {item.text} className = "text-aqua font-bold text-xl"/>
            <Text text = {item.percentage} className = "text-gold"/>
        </div>
    {/each}
</section>
