<script>
	import Button from "../../../components/resuable/Button.svelte";
	import InputField from "../../../components/resuable/InputField.svelte";
</script>



<section class="flex flex-col items-center">
    <div class="w-32 h-32 rounded-full overflow-hidden shadow-md shadow-black border-4 border-white mt-20">
        <img src="./img1.jpg" alt="the logo of the doctor">
    </div>
        <form action="" class="flex flex-col mt-10">
        <InputField type = "text" placeholder = "Username" className = "w-72 h-10 opacity-45"/> 
        <InputField type = "email" placeholder = "Email" className = "w-72 h-10 opacity-45 mt-1"/>
        <InputField type = "password" placeholder = "Password" className = "w-72 h-10 opacity-45 mt-1"/>
        <Button text = "Register" className = "bg-purple w-72 h-10 mt-1"/>
    </form>
</section>