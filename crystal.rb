def crystal(num, cnum = 0, onum = 0, sym = "+")
    if onum == 0
        onum = num
    end
    if onum > cnum
        num  -= 1
        cnum += 1
    else
        if cnum != 0
            onum = -1
            num  += 1
            cnum -= 1
        else
            gets
            exit
        end
    end
    for space in 1..num
        print " "
    end
    for cross in 1..cnum
        print sym
    end
    for cross in 1..cnum
        print sym
    end
    for space in 1..num
        print " "
    end
    print "\n"
    crystal(num, cnum, onum)
end

print "We will perform some magic! Please enter your number of choice: "
inp = gets.to_i
crystal(num=inp)
